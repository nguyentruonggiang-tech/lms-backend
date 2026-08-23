package di

import (
	"lms-worker/ent"
	"lms-worker/internal/common/elastic"
	"lms-worker/internal/common/rabbitmq"
)

func Injection(entClient *ent.Client, elasticClient *elastic.ElasticClient, rabbitmq *rabbitmq.RabbitMQ) {
}
