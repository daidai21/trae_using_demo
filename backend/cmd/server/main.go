package main

import (
	"context"
	"log"

	"ecommerce/internal/handler"
	"ecommerce/internal/middleware"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"ecommerce/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	log.Println("Starting e-commerce server...")

	err := model.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = model.AutoMigrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	userRepo := repository.NewUserRepository(model.DB)
	merchantRepo := repository.NewMerchantRepository(model.DB)
	productRepo := repository.NewProductRepository(model.DB)
	cartRepo := repository.NewCartRepository(model.DB)
	orderRepo := repository.NewOrderRepository(model.DB)

	authService := service.NewAuthService(userRepo)
	merchantService := service.NewMerchantService(merchantRepo)
	productService := service.NewProductService(productRepo, merchantRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	orderService := service.NewOrderService(orderRepo, cartRepo, productRepo, model.DB)

	authHandler := handler.NewAuthHandler(authService)
	merchantHandler := handler.NewMerchantHandler(merchantService)
	productHandler := handler.NewProductHandler(productService)
	cartHandler := handler.NewCartHandler(cartService)
	orderHandler := handler.NewOrderHandler(orderService)

	h := server.Default(server.WithHostPorts(":8080"))

	h.Use(middleware.Recovery())
	h.Use(middleware.CORS())
	h.Use(middleware.ErrorHandler())

	h.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		response.Success(c, map[string]string{
			"status": "ok",
		})
	})

	h.GET("/error", func(ctx context.Context, c *app.RequestContext) {
		response.Error(c, "test error")
	})

	api := h.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProductByID)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			merchant := protected.Group("/merchants")
			{
				merchant.POST("", merchantHandler.CreateMerchant)
				merchant.GET("", merchantHandler.GetMerchants)
				merchant.GET("/:id", merchantHandler.GetMerchantByID)
				merchant.PUT("/:id", merchantHandler.UpdateMerchant)
			}

			products := protected.Group("/products")
			{
				products.POST("", productHandler.CreateProduct)
				products.PUT("/:id", productHandler.UpdateProduct)
				products.DELETE("/:id", productHandler.DeleteProduct)
			}

			cart := protected.Group("/cart")
			{
				cart.POST("", cartHandler.AddToCart)
				cart.GET("", cartHandler.GetCart)
				cart.PUT("/:id", cartHandler.UpdateCartItem)
				cart.DELETE("/:id", cartHandler.DeleteCartItem)
			}

			orders := protected.Group("/orders")
			{
				orders.POST("", orderHandler.CreateOrder)
				orders.GET("", orderHandler.GetOrders)
				orders.GET("/:id", orderHandler.GetOrderByID)
				orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
			}

			protected.GET("/test", func(ctx context.Context, c *app.RequestContext) {
				userID, _ := c.Get("user_id")
				username, _ := c.Get("username")
				response.Success(c, map[string]interface{}{
					"message":  "Authenticated successfully",
					"user_id":  userID,
					"username": username,
				})
			})
		}
	}

	log.Println("Server is running on http://localhost:8080")
	h.Spin()
}
