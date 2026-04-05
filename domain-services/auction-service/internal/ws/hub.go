package ws

import (
	"encoding/json"
	"log"
	"sync"
	"github.com/gorilla/websocket"
)

type MessageType string

const (
	MessageTypeBid      MessageType = "bid"
	MessageTypeUpdate   MessageType = "update"
	MessageTypeJoin     MessageType = "join"
	MessageTypeLeave    MessageType = "leave"
)

type Message struct {
	Type      MessageType `json:"type"`
	AuctionID uint        `json:"auction_id"`
	UserID    uint        `json:"user_id,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan []byte
	AuctionID uint
	UserID    uint
}

type Hub struct {
	clients    map[uint]map[*Client]bool
	Broadcast  chan *Message
	Register   chan *Client
	Unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uint]map[*Client]bool),
		Broadcast:  make(chan *Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			if h.clients[client.AuctionID] == nil {
				h.clients[client.AuctionID] = make(map[*Client]bool)
			}
			h.clients[client.AuctionID][client] = true
			h.mu.Unlock()
			log.Printf("Client %d joined auction %d", client.UserID, client.AuctionID)

		case client := <-h.Unregister:
			h.mu.Lock()
			if clients, ok := h.clients[client.AuctionID]; ok {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					close(client.Send)
					if len(clients) == 0 {
						delete(h.clients, client.AuctionID)
					}
				}
			}
			h.mu.Unlock()
			log.Printf("Client %d left auction %d", client.UserID, client.AuctionID)

		case message := <-h.Broadcast:
			h.mu.RLock()
			if clients, ok := h.clients[message.AuctionID]; ok {
				msgBytes, err := json.Marshal(message)
				if err != nil {
					log.Printf("Error marshaling message: %v", err)
					continue
				}
				for client := range clients {
					select {
					case client.Send <- msgBytes:
					default:
						close(client.Send)
						delete(clients, client)
						if len(clients) == 0 {
							delete(h.clients, message.AuctionID)
						}
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) GetOnlineCount(auctionID uint) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if clients, ok := h.clients[auctionID]; ok {
		return len(clients)
	}
	return 0
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}
		log.Printf("Received message from user %d: %s", c.UserID, message)
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
