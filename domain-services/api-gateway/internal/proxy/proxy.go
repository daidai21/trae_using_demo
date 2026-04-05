package proxy

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Proxy struct {
	upstream string
}

func NewProxy(upstream string) *Proxy {
	return &Proxy{upstream: upstream}
}

func (p *Proxy) Forward(ctx context.Context, c *app.RequestContext) {
	reqBody, err := c.Body()
	if err != nil {
		c.String(consts.StatusInternalServerError, "Failed to read request body")
		return
	}

	method := strings.ToUpper(string(c.Method()))
	path := string(c.Path())
	query := c.Request.URI().QueryString()

	targetURL := p.upstream + path
	if len(query) > 0 {
		targetURL += "?" + string(query)
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, targetURL, bytes.NewReader(reqBody))
	if err != nil {
		c.String(consts.StatusInternalServerError, "Failed to create request")
		return
	}

	c.Request.Header.VisitAll(func(key, value []byte) {
		req.Header.Add(string(key), string(value))
	})

	resp, err := client.Do(req)
	if err != nil {
		c.String(consts.StatusBadGateway, "Failed to reach upstream")
		return
	}
	defer resp.Body.Close()

	c.Response.Header.SetStatusCode(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			c.Response.Header.Add(key, value)
		}
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(consts.StatusInternalServerError, "Failed to read response")
		return
	}

	c.Response.Header.SetContentLength(len(respBody))
	c.Response.SetBodyRaw(respBody)
}

func HealthCheck(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status":  "ok",
		"service": "api-gateway",
		"routes": map[string]string{
			"/api/auth/*":  "user-service:8081",
			"/api/users/*": "user-service:8081",
			"/api/merchants/*": "product-service:8082",
			"/api/products/*":  "product-service:8082",
			"/api/cart/*":      "trade-service:8083",
			"/api/buy/*":       "trade-service:8083",
			"/api/orders/*":    "trade-service:8083",
		},
	})
}
