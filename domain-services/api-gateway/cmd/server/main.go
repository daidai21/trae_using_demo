package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ecommerce/api-gateway/internal/proxy"
	"ecommerce/common/pkg/identity"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	userProxy := proxy.NewProxy("http://localhost:8081")
	productProxy := proxy.NewProxy("http://localhost:8082")
	tradeProxy := proxy.NewProxy("http://localhost:8083")
	auctionProxy := proxy.NewProxy("http://localhost:8084")

	platformUserProxy := proxy.NewProxy("http://localhost:9081")
	platformProductProxy := proxy.NewProxy("http://localhost:9082")
	platformTradeProxy := proxy.NewProxy("http://localhost:9083")

	resolver := identity.NewDefaultIdentityResolver()

	h := server.Default(server.WithHostPorts(":8080"))

	h.Use(func(ctx context.Context, c *app.RequestContext) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Business-Identity")
		if string(c.Method()) == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		req := &http.Request{
			Header: make(http.Header),
			URL:    c.Request.URI(),
		}
		c.Request.Header.VisitAll(func(key, value []byte) {
			req.Header.Add(string(key), string(value))
		})

		bizIdentity, _ := resolver.Resolve(ctx, req)
		if bizIdentity != nil {
			c.Set("business_identity", bizIdentity)
			log.Printf("Business identity: %s", bizIdentity.String())
		}

		c.Next(ctx)
	})

	h.GET("/health", proxy.HealthCheck)

	api := h.Group("/api")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.Any("/*path", userProxy.Forward)
		}

		usersGroup := api.Group("/users")
		{
			usersGroup.Any("/*path", userProxy.Forward)
		}

		merchantsGroup := api.Group("/merchants")
		{
			merchantsGroup.Any("/*path", productProxy.Forward)
		}

		productsGroup := api.Group("/products")
		{
			productsGroup.Any("/*path", productProxy.Forward)
		}

		cartGroup := api.Group("/cart")
		{
			cartGroup.Any("/*path", tradeProxy.Forward)
		}

		buyGroup := api.Group("/buy")
		{
			buyGroup.Any("/*path", tradeProxy.Forward)
		}

		ordersGroup := api.Group("/orders")
		{
			ordersGroup.Any("/*path", tradeProxy.Forward)
		}

		auctionsGroup := api.Group("/auctions")
		{
			auctionsGroup.Any("/*path", auctionProxy.Forward)
		}
	}

	ws := h.Group("/api/ws")
	{
		ws.Any("/*path", auctionProxy.Forward)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down API gateway...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := h.Shutdown(ctx); err != nil {
			log.Fatal("API gateway shutdown error:", err)
		}
	}()

	log.Println("API gateway (platform v4.0) starting on :8080")
	h.Spin()
}
