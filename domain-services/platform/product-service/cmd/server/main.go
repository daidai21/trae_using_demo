package main

import (
	"log"

	"ecommerce/product-service/internal/handler"
	"ecommerce/product-service/internal/middleware"
	"ecommerce/product-service/internal/model"
	"ecommerce/product-service/internal/repository"
	"ecommerce/product-service/internal/service"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	db, err := model.InitDB("product.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	merchantRepo := repository.NewMerchantRepository(db)
	productRepo := repository.NewProductRepository(db)
	merchantService := service.NewMerchantService(merchantRepo)
	productService := service.NewProductService(productRepo, merchantRepo)
	merchantHandler := handler.NewMerchantHandler(merchantService)
	productHandler := handler.NewProductHandler(productService)

	h := server.Default(server.WithHostPorts(":8082"))

	h.POST("/api/merchants", middleware.AuthMiddleware(), merchantHandler.CreateMerchant)
	h.GET("/api/merchants", merchantHandler.GetMerchants)
	h.GET("/api/merchants/:id", merchantHandler.GetMerchantByID)
	h.PUT("/api/merchants/:id", middleware.AuthMiddleware(), merchantHandler.UpdateMerchant)

	h.POST("/api/products", middleware.AuthMiddleware(), productHandler.CreateProduct)
	h.GET("/api/products", productHandler.GetProducts)
	h.GET("/api/products/:id", productHandler.GetProductByID)
	h.PUT("/api/products/:id", middleware.AuthMiddleware(), productHandler.UpdateProduct)
	h.DELETE("/api/products/:id", middleware.AuthMiddleware(), productHandler.DeleteProduct)

	log.Println("Product Service starting on :8082")
	h.Spin()
}
