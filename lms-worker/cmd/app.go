package main

import (
	"lms-worker/ent"
	"lms-worker/internal/common/elastic"
	"lms-worker/internal/common/ent_client"
	"lms-worker/internal/common/env"
	"lms-worker/internal/common/rabbitmq"
	"lms-worker/internal/di"
)

type App struct {
	env       *env.Env
	entClient *ent.Client
	rabbitmq  *rabbitmq.RabbitMQ
}

func NewApp() *App {
	e := env.New()
	ec := ent_client.New(e)
	elasticClient := elastic.NewElasticClient(e)
	rmq := rabbitmq.NewRabbitMQ(e)
	di.Injection(ec, elasticClient, rmq)

	return &App{env: e, entClient: ec, rabbitmq: rmq}
}

func (a *App) Start() {
	select {}
}
