package delivery

import (
	"lms-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type categoryDelivery struct {
	categoryHandler *handler.CategoryHandler
}

func NewCategoryDelivery(categoryHandler *handler.CategoryHandler) *categoryDelivery {
	return &categoryDelivery{categoryHandler: categoryHandler}
}

func (d *categoryDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	apiGroup.GET("categories", d.categoryHandler.FindAll)
}
