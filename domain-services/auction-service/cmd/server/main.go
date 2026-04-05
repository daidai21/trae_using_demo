package main

import (
	"ecommerce/auction-service/internal/handler"
	"ecommerce/auction-service/internal/middleware"
	"ecommerce/auction-service/internal/model"
	"ecommerce/auction-service/internal/repository"
	"ecommerce/auction-service/internal/service"
	"ecommerce/auction-service/internal/ws"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	db, err := model.InitDB("./auction.db")
	if err != nil {
		log.Fatal("Failed to init database:", err)
	}

	hub := ws.NewHub()
	go hub.Run()

	auctionRepo := repository.NewAuctionRepository(db)
	bidRepo := repository.NewBidRepository(db)
	auctionService := service.NewAuctionService(auctionRepo, bidRepo, hub)
	auctionHandler := handler.NewAuctionHandler(auctionService, hub)

	h := server.Default(server.WithHostPorts(":8084"))

	auth := h.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/auctions", auctionHandler.CreateAuction)
		auth.POST("/auctions/:id/start", auctionHandler.StartAuction)
		auth.POST("/auctions/:id/end", auctionHandler.EndAuction)
		auth.POST("/auctions/:id/bid", auctionHandler.PlaceBid)
	}

	public := h.Group("/api")
	{
		public.GET("/auctions", auctionHandler.GetAllAuctions)
		public.GET("/auctions/live", auctionHandler.GetLiveAuctions)
		public.GET("/auctions/:id", auctionHandler.GetAuction)
		public.GET("/ws/auctions/:id", auctionHandler.WebSocket)
	}

	log.Println("Auction service starting on :8084")
	h.Spin()
}
