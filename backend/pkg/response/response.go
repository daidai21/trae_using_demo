package response

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	CodeSuccess = 0
	CodeError   = 1
)

func Success(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

func Error(c *app.RequestContext, message string) {
	c.JSON(consts.StatusOK, Response{
		Code:    CodeError,
		Message: message,
	})
}

func ErrorWithCode(c *app.RequestContext, code int, message string) {
	c.JSON(consts.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

func ErrorWithStatus(c *app.RequestContext, httpStatus int, message string) {
	c.JSON(httpStatus, Response{
		Code:    CodeError,
		Message: message,
	})
}

func BadRequest(c *app.RequestContext, message string) {
	ErrorWithStatus(c, http.StatusBadRequest, message)
}

func Unauthorized(c *app.RequestContext, message string) {
	ErrorWithStatus(c, http.StatusUnauthorized, message)
}

func Forbidden(c *app.RequestContext, message string) {
	ErrorWithStatus(c, http.StatusForbidden, message)
}

func NotFound(c *app.RequestContext, message string) {
	ErrorWithStatus(c, http.StatusNotFound, message)
}

func InternalServerError(c *app.RequestContext, message string) {
	ErrorWithStatus(c, http.StatusInternalServerError, message)
}
