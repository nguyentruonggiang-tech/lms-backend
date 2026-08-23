package delivery

import (
	"context"
	"fmt"

	"lms-worker/internal/common/rabbitmq"
	"lms-worker/internal/handler"
)

type SearchDelivery struct {
	rabbitmq      *rabbitmq.RabbitMQ
	searchHandler *handler.SearchHandler
}

func NewSearchDelivery(rabbitmq *rabbitmq.RabbitMQ, searchHandler *handler.SearchHandler) *SearchDelivery {
	return &SearchDelivery{rabbitmq: rabbitmq, searchHandler: searchHandler}
}

func (d *SearchDelivery) Register(ctx context.Context) {
	if err := d.rabbitmq.On(ctx, "search.course_index", d.searchHandler.HandleCourseIndex); err != nil {
		fmt.Printf("❌ [DELIVERY] failed to register search.course_index: %v\n", err)
	}
}
