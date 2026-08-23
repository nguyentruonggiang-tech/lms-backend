package main

import (
	"fmt"

	"lms-api/internal/common/ent_client"
	"lms-api/internal/common/env"
	"lms-api/internal/common/middlewares"
	"lms-api/internal/common/response"
	"lms-api/internal/common/swagger"
	"lms-api/internal/di"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	ginEngine *gin.Engine
	env       *env.Env
}

func NewApp() *App {
	e := env.New()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middlewares.ErrorHandler)
	r.Use(gin.CustomRecovery(func(ctx *gin.Context, _ any) {
		ctx.Error(response.NewInternalServerErrorException())
		ctx.Abort()
	}))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	entClient := ent_client.New(e)
	di.Injection(r, entClient, e)
	swagger.Start(r)

	return &App{ginEngine: r, env: e}
}

func (a *App) Start() {
	addr := fmt.Sprintf("%s:%s", a.env.Host, a.env.Port)
	fmt.Printf("🚀 Server running at http://%s\n", addr)
	a.ginEngine.Run(addr)
}
