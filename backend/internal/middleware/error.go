package middleware

import (
	"context"
	"ecommerce/pkg/response"
	"net/http"
	"runtime/debug"

	"github.com/cloudwego/hertz/pkg/app"
)

func Recovery() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				response.InternalServerError(c, "internal server error")
				c.Abort()
			}
		}()
		c.Next(ctx)
	}
}

func ErrorHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			switch c.Response.StatusCode() {
			case http.StatusBadRequest:
				response.BadRequest(c, err.Error())
			case http.StatusUnauthorized:
				response.Unauthorized(c, err.Error())
			case http.StatusForbidden:
				response.Forbidden(c, err.Error())
			case http.StatusNotFound:
				response.NotFound(c, err.Error())
			default:
				response.InternalServerError(c, err.Error())
			}
		}
	}
}
