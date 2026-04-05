package handler

import (
	"context"
	"strconv"

	"ecommerce/internal/service"
	"ecommerce/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	var req service.CreateProductRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	product, err := h.productService.CreateProduct(userID.(uint), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, product)
}

func (h *ProductHandler) GetProducts(ctx context.Context, c *app.RequestContext) {
	products, err := h.productService.GetProducts()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, products)
}

func (h *ProductHandler) GetProductByID(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		response.NotFound(c, "product not found")
		return
	}

	response.Success(c, product)
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	var req service.UpdateProductRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	product, err := h.productService.UpdateProduct(userID.(uint), uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, product)
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	if err := h.productService.DeleteProduct(userID.(uint), uint(id)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, map[string]string{
		"message": "product deleted successfully",
	})
}
