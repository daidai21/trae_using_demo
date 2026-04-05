package handler

import (
	"context"
	"encoding/json"

	"ecommerce/common/pkg/response"
	"ecommerce/trade-service/internal/model"
	"ecommerce/trade-service/internal/service/buy"

	"github.com/cloudwego/hertz/pkg/app"
)

type BuyHandler struct {
	buyService  *buy.BuyService
	cartService interface {
		GetCart(userID uint) ([]*model.Cart, error)
	}
}

func NewBuyHandler(buyService *buy.BuyService, cartService interface {
	GetCart(userID uint) ([]*model.Cart, error)
}) *BuyHandler {
	return &BuyHandler{buyService: buyService, cartService: cartService}
}

func (h *BuyHandler) CreateOrder(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	carts, err := h.cartService.GetCart(userID.(uint))
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	productsJson := string(c.GetHeader("X-Products"))
	products := make(map[uint]*buy.ProductInfo)
	if productsJson != "" {
		var productsList []map[string]interface{}
		if err := json.Unmarshal([]byte(productsJson), &productsList); err == nil {
			for _, p := range productsList {
				id := uint(p["id"].(float64))
				products[id] = &buy.ProductInfo{
					ID:    id,
					Price: p["price"].(float64),
					Stock: int(p["stock"].(float64)),
					Name:  p["name"].(string),
				}
			}
		}
	}

	if len(products) == 0 {
		for _, cart := range carts {
			products[cart.ProductID] = &buy.ProductInfo{
				ID:    cart.ProductID,
				Price: 100.0,
				Stock: 9999,
				Name:  "Product",
			}
		}
	}

	order, err := h.buyService.CreateOrder(userID.(uint), carts, products)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, order)
}
