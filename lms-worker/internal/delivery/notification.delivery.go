package delivery

import (
	"context"
	"fmt"

	"lms-worker/internal/common/rabbitmq"
	"lms-worker/internal/handler"
)

type NotificationDelivery struct {
	rabbitmq            *rabbitmq.RabbitMQ
	notificationHandler *handler.NotificationHandler
}

func NewNotificationDelivery(rabbitmq *rabbitmq.RabbitMQ, notificationHandler *handler.NotificationHandler) *NotificationDelivery {
	return &NotificationDelivery{rabbitmq: rabbitmq, notificationHandler: notificationHandler}
}

func (d *NotificationDelivery) Register(ctx context.Context) {
	if err := d.rabbitmq.On(ctx, "notification.course_enrolled", d.notificationHandler.HandleCourseEnrolled); err != nil {
		fmt.Printf("❌ [DELIVERY] failed to register notification.course_enrolled: %v\n", err)
	}
}
