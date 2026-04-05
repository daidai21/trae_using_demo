import { useEffect, useRef, useState, useCallback } from 'react'

export interface AuctionMessage {
  type: 'bid' | 'update' | 'join' | 'leave'
  auction_id: number
  user_id?: number
  data?: any
  timestamp: number
}

export const useAuctionSocket = (auctionId: number, userId?: number) => {
  const [isConnected, setIsConnected] = useState(false)
  const [messages, setMessages] = useState<AuctionMessage[]>([])
  const [onlineCount, setOnlineCount] = useState(0)
  const wsRef = useRef<WebSocket | null>(null)

  const connect = useCallback(() => {
    if (wsRef.current?.readyState === WebSocket.OPEN) return

    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
    const host = window.location.host
    let url = `${protocol}//${host}/api/ws/auctions/${auctionId}`
    if (userId) {
      url += `?user_id=${userId}`
    }

    const ws = new WebSocket(url)
    wsRef.current = ws

    ws.onopen = () => {
      setIsConnected(true)
    }

    ws.onmessage = (event) => {
      try {
        const message: AuctionMessage = JSON.parse(event.data)
        setMessages((prev) => [...prev, message])
        
        if (message.data?.online_count !== undefined) {
          setOnlineCount(message.data.online_count)
        }
      } catch (error) {
        console.error('Failed to parse message:', error)
      }
    }

    ws.onclose = () => {
      setIsConnected(false)
    }

    ws.onerror = (error) => {
      console.error('WebSocket error:', error)
    }
  }, [auctionId, userId])

  const disconnect = useCallback(() => {
    if (wsRef.current) {
      wsRef.current.close()
      wsRef.current = null
      setIsConnected(false)
    }
  }, [])

  useEffect(() => {
    connect()
    return () => {
      disconnect()
    }
  }, [connect, disconnect])

  return {
    isConnected,
    messages,
    onlineCount,
    connect,
    disconnect,
  }
}
