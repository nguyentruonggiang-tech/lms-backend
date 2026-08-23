package di

import (
	"lms-worker/ent"
	"lms-worker/internal/common/elastic"
	"lms-worker/internal/common/rabbitmq"
	"lms-worker/internal/delivery"
	"lms-worker/internal/handler"
	"lms-worker/internal/repository/repository_impl"
	"lms-worker/internal/usecase/usecase_impl"
)

func Injection(entClient *ent.Client, elasticClient *elastic.ElasticClient, rabbitmq *rabbitmq.RabbitMQ) *delivery.RootDelivery {
	notificationRepository := repository_impl.NewNotificationRepository(entClient)
	notificationUsecase := usecase_impl.NewNotificationUsecase(notificationRepository)
	notificationHandler := handler.NewNotificationHandler(notificationUsecase)
	notificationDelivery := delivery.NewNotificationDelivery(rabbitmq, notificationHandler)

	searchUsecase := usecase_impl.NewSearchUsecase(elasticClient)
	searchHandler := handler.NewSearchHandler(searchUsecase)
	searchDelivery := delivery.NewSearchDelivery(rabbitmq, searchHandler)

	return delivery.NewRootDelivery(notificationDelivery, searchDelivery)
}
