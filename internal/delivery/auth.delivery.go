package delivery

import (
	"lms-backend/internal/common/middlewares"
	"lms-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type authDelivery struct {
	handler        *handler.AuthHandler
	authMiddleware *middlewares.AuthMiddleware
}

func NewAuthDelivery(h *handler.AuthHandler, m *middlewares.AuthMiddleware) *authDelivery {
	return &authDelivery{handler: h, authMiddleware: m}
}

func (d *authDelivery) RegisterRouter(api *gin.RouterGroup) {
	authGroup := api.Group("auth")
	authGroup.POST("register", d.handler.Register)
	authGroup.POST("login", d.handler.Login)
	authGroup.POST("refresh-token", d.handler.RefreshToken)

	protected := authGroup.Group("")
	protected.Use(d.authMiddleware.Protect)
	protected.GET("get-info", d.handler.GetInfo)
	protected.PATCH("change-password", d.handler.ChangePassword)
}
