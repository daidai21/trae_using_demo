package handler

import (
	"context"
	"strconv"

	"ecommerce/common/pkg/response"
	"ecommerce/trade-service/internal/service/cart"

	"github.com/cloudwego/hertz/pkg/app"
)

type CartHandler struct {
	cartService *cart.CartService
}

func NewCartHandler(cartService *cart.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) AddToCart(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	var req cart.AddToCartRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	productStock, _ := strconv.Atoi(string(c.Query("product_stock")))
	if productStock == 0 {
		productStock = 99999
	}

	cartItem, err := h.cartService.AddToCart(userID.(uint), req.ProductID, req.Quantity, productStock)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, cartItem)
}

func (h *CartHandler) GetCart(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	carts, err := h.cartService.GetCart(userID.(uint))
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, carts)
}

func (h *CartHandler) UpdateCartItem(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid cart id")
		return
	}

	var req cart.UpdateCartRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	productStock, _ := strconv.Atoi(string(c.Query("product_stock")))
	if productStock == 0 {
		productStock = 99999
	}

	cartItem, err := h.cartService.UpdateCartItem(userID.(uint), uint(id), req.Quantity, productStock)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, cartItem)
}

func (h *CartHandler) DeleteCartItem(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid cart id")
		return
	}

	if err := h.cartService.DeleteCartItem(userID.(uint), uint(id)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, map[string]string{
		"message": "cart item deleted successfully",
	})
}
