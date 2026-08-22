package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type successRes struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type errorRes struct {
	Message string `json:"message"`
}

func Success(data any, message string, code int, ctx *gin.Context) {
	if code == 0 {
		code = http.StatusOK
	}
	if message == "" {
		message = "success"
	}
	ctx.JSON(code, successRes{Data: data, Message: message})
}

func Error(message string, code int, ctx *gin.Context) {
	ctx.JSON(code, errorRes{Message: message})
}
