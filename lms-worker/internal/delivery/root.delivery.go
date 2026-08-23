package delivery

import "context"

type RootDelivery struct {
	notificationDelivery *NotificationDelivery
}

func NewRootDelivery(notificationDelivery *NotificationDelivery) *RootDelivery {
	return &RootDelivery{notificationDelivery: notificationDelivery}
}

func (r *RootDelivery) Start(ctx context.Context) {
	r.notificationDelivery.Register(ctx)
}
