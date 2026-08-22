package middlewares

import (
	"errors"
	"net/http"

	"lms-backend/internal/common/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err
		code := http.StatusInternalServerError
		msg := http.StatusText(code)

		var ex *response.Exception
		if errors.As(err, &ex) {
			code = ex.StatusCode
			msg = ex.Message
		}

		response.Error(msg, code, c)
	}
}
