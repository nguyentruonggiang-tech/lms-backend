package main

import (
	"context"

	"lms-worker/ent"
	"lms-worker/internal/common/elastic"
	"lms-worker/internal/common/ent_client"
	"lms-worker/internal/common/env"
	"lms-worker/internal/common/rabbitmq"
	"lms-worker/internal/delivery"
	"lms-worker/internal/di"
)

type App struct {
	env          *env.Env
	entClient    *ent.Client
	rabbitmq     *rabbitmq.RabbitMQ
	rootDelivery *delivery.RootDelivery
}

func NewApp() *App {
	e := env.New()
	ec := ent_client.New(e)
	elasticClient := elastic.NewElasticClient(e)
	rmq := rabbitmq.NewRabbitMQ(e)
	rootDelivery := di.Injection(ec, elasticClient, rmq)

	return &App{env: e, entClient: ec, rabbitmq: rmq, rootDelivery: rootDelivery}
}

func (a *App) Start() {
	ctx := context.Background()
	a.rootDelivery.Start(ctx)
	select {}
}
