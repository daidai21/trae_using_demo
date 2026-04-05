package handler

import (
	"context"
	"strconv"

	"ecommerce/common/pkg/response"
	"ecommerce/product-service/internal/service"

	"github.com/cloudwego/hertz/pkg/app"
)

type MerchantHandler struct {
	merchantService *service.MerchantService
}

func NewMerchantHandler(merchantService *service.MerchantService) *MerchantHandler {
	return &MerchantHandler{merchantService: merchantService}
}

func (h *MerchantHandler) CreateMerchant(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	var req service.CreateMerchantRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	merchant, err := h.merchantService.CreateMerchant(userID.(uint), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, merchant)
}

func (h *MerchantHandler) GetMerchants(ctx context.Context, c *app.RequestContext) {
	merchants, err := h.merchantService.GetMerchants()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, merchants)
}

func (h *MerchantHandler) GetMerchantByID(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid merchant id")
		return
	}

	merchant, err := h.merchantService.GetMerchantByID(uint(id))
	if err != nil {
		response.NotFound(c, "merchant not found")
		return
	}

	response.Success(c, merchant)
}

func (h *MerchantHandler) UpdateMerchant(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid merchant id")
		return
	}

	var req service.UpdateMerchantRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	merchant, err := h.merchantService.UpdateMerchant(userID.(uint), uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, merchant)
}
