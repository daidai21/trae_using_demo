package handler

import (
	"context"
	"strconv"

	"ecommerce/internal/service"
	"ecommerce/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	order, err := h.orderService.CreateOrder(userID.(uint))
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, order)
}

func (h *OrderHandler) GetOrders(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	orders, err := h.orderService.GetOrders(userID.(uint))
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, orders)
}

func (h *OrderHandler) GetOrderByID(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid order id")
		return
	}

	order, err := h.orderService.GetOrderByID(userID.(uint), uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, order)
}

func (h *OrderHandler) UpdateOrderStatus(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid order id")
		return
	}

	var req service.UpdateOrderStatusRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	order, err := h.orderService.UpdateOrderStatus(userID.(uint), uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, order)
}
