package di

import (
	"lms-backend/ent"
	"lms-backend/internal/common/env"
	"lms-backend/internal/common/middlewares"
	"lms-backend/internal/delivery"
	"lms-backend/internal/handler"
	"lms-backend/internal/repository/repository_impl"
	"lms-backend/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, e *env.Env) {
	tokenUsecase := usecase_impl.NewTokenUsecase(e)
	userRepo := repository_impl.NewUserRepository(entClient)
	authMiddleware := middlewares.NewAuthMiddleware(tokenUsecase, userRepo)

	authUsecase := usecase_impl.NewAuthUsecase(userRepo, tokenUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)
	authDelivery := delivery.NewAuthDelivery(authHandler, authMiddleware)

	rootDelivery := delivery.NewRootDelivery(authDelivery)
	rootDelivery.RegisterRouter(ginEngine)
}
