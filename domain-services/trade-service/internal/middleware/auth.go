package middleware

import (
	"context"
	"ecommerce/common/pkg/utils"
	"ecommerce/common/pkg/response"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			response.Unauthorized(c, "authorization header required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "invalid authorization header format")
			c.Abort()
			return
		}

		userID, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next(ctx)
	}
}
