package main

import (
	"context"
	"log"

	"ecommerce/common/pkg/response"
	"ecommerce/user-service/internal/handler"
	"ecommerce/user-service/internal/middleware"
	"ecommerce/user-service/internal/model"
	"ecommerce/user-service/internal/repository"
	"ecommerce/user-service/internal/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	db, err := model.InitDB("users.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	h := server.Default(server.WithHostPorts(":8081"))

	h.POST("/api/auth/register", authHandler.Register)
	h.POST("/api/auth/login", authHandler.Login)

	h.GET("/api/users/profile", middleware.AuthMiddleware(), func(ctx context.Context, c *app.RequestContext) {
		userID, _ := c.Get("user_id")
		username, _ := c.Get("username")
		response.Success(c, map[string]interface{}{
			"id":       userID,
			"username": username,
		})
	})

	log.Println("User Service starting on :8081")
	h.Spin()
}
