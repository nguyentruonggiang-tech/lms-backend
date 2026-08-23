package delivery

import "context"

type RootDelivery struct {
	notificationDelivery *NotificationDelivery
	searchDelivery       *SearchDelivery
}

func NewRootDelivery(notificationDelivery *NotificationDelivery, searchDelivery *SearchDelivery) *RootDelivery {
	return &RootDelivery{
		notificationDelivery: notificationDelivery,
		searchDelivery:       searchDelivery,
	}
}

func (r *RootDelivery) Start(ctx context.Context) {
	r.notificationDelivery.Register(ctx)
	r.searchDelivery.Register(ctx)
}
