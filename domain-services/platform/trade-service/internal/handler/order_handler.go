package handler

import (
	"context"
	"strconv"

	"ecommerce/common/pkg/response"
	"ecommerce/trade-service/internal/service/order"

	"github.com/cloudwego/hertz/pkg/app"
)

type OrderHandler struct {
	orderService *order.OrderService
}

func NewOrderHandler(orderService *order.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
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

	orderItem, err := h.orderService.GetOrderByID(userID.(uint), uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, orderItem)
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

	var req order.UpdateOrderStatusRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	orderItem, err := h.orderService.UpdateOrderStatus(userID.(uint), uint(id), req.Status)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, orderItem)
}
