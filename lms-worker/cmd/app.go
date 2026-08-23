package main

import (
	"lms-worker/internal/common/ent_client"
	"lms-worker/internal/common/env"
	"lms-worker/internal/di"

	"lms-worker/ent"
)

type App struct {
	env       *env.Env
	entClient *ent.Client
}

func NewApp() *App {
	e := env.New()
	ec := ent_client.New(e)
	di.Injection(ec)

	return &App{env: e, entClient: ec}
}

func (a *App) Start() {
	select {}
}
