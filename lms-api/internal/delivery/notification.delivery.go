package delivery

import (
	"lms-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type notificationDelivery struct {
	notificationHandler *handler.NotificationHandler
}

func NewNotificationDelivery(notificationHandler *handler.NotificationHandler) *notificationDelivery {
	return &notificationDelivery{notificationHandler: notificationHandler}
}

func (d *notificationDelivery) RegisterRouter(studentGroup *gin.RouterGroup) {
	myGroup := studentGroup.Group("my")
	{
		myGroup.GET("notifications", d.notificationHandler.GetMyNotifications)
		myGroup.PATCH("notifications/:id/read", d.notificationHandler.MarkRead)
		myGroup.PATCH("notifications/read-all", d.notificationHandler.MarkAllRead)
	}
}
