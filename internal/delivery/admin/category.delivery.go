package admin_delivery

import (
	"lms-backend/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type CategoryDelivery struct {
	categoryHandler *admin.CategoryHandler
}

func NewCategoryDelivery(categoryHandler *admin.CategoryHandler) *CategoryDelivery {
	return &CategoryDelivery{
		categoryHandler: categoryHandler,
	}
}

func (d *CategoryDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	categories := adminGroup.Group("categories")
	{
		categories.POST("", d.categoryHandler.Create)
		categories.GET("", d.categoryHandler.FindAll)
		categories.GET(":id", d.categoryHandler.FindByID)
		categories.PUT(":id", d.categoryHandler.Update)
		categories.DELETE(":id", d.categoryHandler.Delete)
	}
}
