package delivery

import (
	"lms-backend/internal/common/middlewares"
	"lms-backend/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	authDelivery     *authDelivery
	categoryDelivery *admin_delivery.CategoryDelivery
	authMiddleware   *middlewares.AuthMiddleware
}

func NewRootDelivery(
	authDelivery *authDelivery,
	categoryDelivery *admin_delivery.CategoryDelivery,
	authMiddleware *middlewares.AuthMiddleware,
) *rootDelivery {
	return &rootDelivery{
		authDelivery:     authDelivery,
		categoryDelivery: categoryDelivery,
		authMiddleware:   authMiddleware,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{
		r.authDelivery.RegisterRouter(apiGroup)

		adminGroup := apiGroup.Group("admin")
		adminGroup.Use(r.authMiddleware.Protect, r.authMiddleware.AdminOnly)
		{
			r.categoryDelivery.RegisterRouter(adminGroup)
		}
	}
}
