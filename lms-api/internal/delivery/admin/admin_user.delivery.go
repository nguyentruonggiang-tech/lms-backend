package admin

import (
	"lms-api/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userHandler *admin.UserHandler
}

func NewUserDelivery(userHandler *admin.UserHandler) *UserDelivery {
	return &UserDelivery{userHandler: userHandler}
}

func (d *UserDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	u := adminGroup.Group("users")
	{
		u.GET("", d.userHandler.FindAll)
		u.GET(":id", d.userHandler.FindByID)
		u.PATCH(":id/status", d.userHandler.UpdateStatus)
		u.PATCH(":id/role", d.userHandler.UpdateRole)
	}
}
