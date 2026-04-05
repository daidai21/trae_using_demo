package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ecommerce/common/pkg/response"
	"ecommerce/trade-service/internal/handler"
	"ecommerce/trade-service/internal/middleware"
	"ecommerce/trade-service/internal/model"
	"ecommerce/trade-service/internal/repository"
	"ecommerce/trade-service/internal/service/buy"
	"ecommerce/trade-service/internal/service/cart"
	"ecommerce/trade-service/internal/service/order"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	db, err := model.InitDB("trade.db")
	if err != nil {
		log.Fatal("Failed to init database:", err)
	}

	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	cartService := cart.NewCartService(cartRepo)
	orderService := order.NewOrderService(orderRepo)
	buyService := buy.NewBuyService(orderRepo, cartRepo, db)

	cartHandler := handler.NewCartHandler(cartService)
	orderHandler := handler.NewOrderHandler(orderService)
	buyHandler := handler.NewBuyHandler(buyService, cartService)

	h := server.Default(server.WithHostPorts(":8083"))

	h.Use(func(ctx context.Context, c *app.RequestContext) {
		c.Header("Content-Type", "application/json")
		c.Next(ctx)
	})

	h.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		response.Success(c, map[string]string{
			"status": "ok",
			"service": "trade-service",
		})
	})

	api := h.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		cartGroup := api.Group("/cart")
		{
			cartGroup.POST("/", cartHandler.AddToCart)
			cartGroup.GET("/", cartHandler.GetCart)
			cartGroup.PUT("/:id", cartHandler.UpdateCartItem)
			cartGroup.DELETE("/:id", cartHandler.DeleteCartItem)
		}

		buyGroup := api.Group("/buy")
		{
			buyGroup.POST("/", buyHandler.CreateOrder)
		}

		orderGroup := api.Group("/orders")
		{
			orderGroup.GET("/", orderHandler.GetOrders)
			orderGroup.GET("/:id", orderHandler.GetOrderByID)
			orderGroup.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := h.Shutdown(ctx); err != nil {
			log.Fatal("Server shutdown error:", err)
		}
	}()

	log.Println("Trade service starting on :8083")
	h.Spin()
}
