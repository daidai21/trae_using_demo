package handler

import (
	"context"
	"strconv"

	"ecommerce/internal/service"
	"ecommerce/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

type CartHandler struct {
	cartService *service.CartService
}

func NewCartHandler(cartService *service.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) AddToCart(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	var req service.AddToCartRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	cart, err := h.cartService.AddToCart(userID.(uint), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, cart)
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

	var req service.UpdateCartRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	cart, err := h.cartService.UpdateCartItem(userID.(uint), uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, cart)
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
