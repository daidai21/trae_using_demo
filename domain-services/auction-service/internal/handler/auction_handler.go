package handler

import (
	"context"
	"ecommerce/auction-service/internal/middleware"
	"ecommerce/auction-service/internal/model"
	"ecommerce/auction-service/internal/service"
	"ecommerce/auction-service/internal/ws"
	"ecommerce/common/pkg/response"
	"net/http"
	"strconv"
	"time"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gorilla/websocket"
)

type AuctionHandler struct {
	auctionService service.AuctionService
	hub            *ws.Hub
	upgrader       websocket.Upgrader
}

func NewAuctionHandler(auctionService service.AuctionService, hub *ws.Hub) *AuctionHandler {
	return &AuctionHandler{
		auctionService: auctionService,
		hub:            hub,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *AuctionHandler) CreateAuction(ctx context.Context, c *app.RequestContext) {
	var auction model.Auction
	if err := c.Bind(&auction); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.Unauthorized(c, "未认证")
		return
	}

	if err := h.auctionService.CreateAuction(&auction, userID); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, auction)
}

func (h *AuctionHandler) GetAuction(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	auction, err := h.auctionService.GetAuction(uint(id))
	if err != nil {
		response.NotFound(c, "拍卖不存在")
		return
	}

	response.Success(c, auction)
}

func (h *AuctionHandler) GetAllAuctions(ctx context.Context, c *app.RequestContext) {
	auctions, err := h.auctionService.GetAllAuctions()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, auctions)
}

func (h *AuctionHandler) GetLiveAuctions(ctx context.Context, c *app.RequestContext) {
	auctions, err := h.auctionService.GetLiveAuctions()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, auctions)
}

func (h *AuctionHandler) StartAuction(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.auctionService.StartAuction(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *AuctionHandler) PlaceBid(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.Bind(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.Unauthorized(c, "未认证")
		return
	}

	bid, err := h.auctionService.PlaceBid(uint(id), userID, req.Amount)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, bid)
}

func (h *AuctionHandler) EndAuction(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	if err := h.auctionService.EndAuction(uint(id)); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *AuctionHandler) WebSocket(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "无效的ID")
		return
	}

	userIDStr := c.Query("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		userID = 0
	}

	conn, err := h.upgrader.Upgrade(&c.Response, &c.Request, nil)
	if err != nil {
		response.InternalServerError(c, "WebSocket连接失败")
		return
	}

	client := &ws.Client{
		Hub:       h.hub,
		Conn:      conn,
		Send:      make(chan []byte, 256),
		AuctionID: uint(id),
		UserID:    uint(userID),
	}

	client.Hub.register <- client

	joinMsg := &ws.Message{
		Type:      ws.MessageTypeJoin,
		AuctionID: uint(id),
		UserID:    uint(userID),
		Data: map[string]interface{}{
			"online_count": h.hub.GetOnlineCount(uint(id)),
		},
		Timestamp: time.Now().Unix(),
	}
	h.hub.Broadcast <- joinMsg

	go client.WritePump()
	client.ReadPump()
}
