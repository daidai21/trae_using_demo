package handler

import (
	"context"

	"ecommerce/common/pkg/response"
	"ecommerce/user-service/internal/service"

	"github.com/cloudwego/hertz/pkg/app"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(ctx context.Context, c *app.RequestContext) {
	var req service.RegisterRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	resp, err := h.authService.Register(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AuthHandler) Login(ctx context.Context, c *app.RequestContext) {
	var req service.LoginRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	resp, err := h.authService.Login(&req)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.Success(c, resp)
}
