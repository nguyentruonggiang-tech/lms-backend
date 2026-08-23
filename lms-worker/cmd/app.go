package main

import (
	"lms-worker/ent"
	"lms-worker/internal/common/elastic"
	"lms-worker/internal/common/ent_client"
	"lms-worker/internal/common/env"
	"lms-worker/internal/di"
)

type App struct {
	env       *env.Env
	entClient *ent.Client
}

func NewApp() *App {
	e := env.New()
	ec := ent_client.New(e)
	elasticClient := elastic.NewElasticClient(e)
	di.Injection(ec, elasticClient)

	return &App{env: e, entClient: ec}
}

func (a *App) Start() {
	select {}
}
