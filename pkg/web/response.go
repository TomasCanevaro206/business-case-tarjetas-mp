package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ResponseMessage(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	ResponseMessage(c, status, Response{Data: data})
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := ErrorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	ResponseMessage(c, status, err)
}
