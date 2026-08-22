package di

import (
	"lms-backend/ent"
	"lms-backend/internal/common/env"
	"lms-backend/internal/common/middlewares"
	"lms-backend/internal/delivery"
	"lms-backend/internal/delivery/admin"
	"lms-backend/internal/handler"
	adminHandler "lms-backend/internal/handler/admin"
	"lms-backend/internal/repository/repository_impl"
	"lms-backend/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, e *env.Env) {
	tokenUsecase := usecase_impl.NewTokenUsecase(e)

	userRepository := repository_impl.NewUserRepository(entClient)
	authMiddleware := middlewares.NewAuthMiddleware(tokenUsecase, userRepository)

	authUsecase := usecase_impl.NewAuthUsecase(userRepository, tokenUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)
	authDelivery := delivery.NewAuthDelivery(authHandler, authMiddleware)

	categoryRepository := repository_impl.NewCategoryRepository(entClient)
	categoryUsecase := usecase_impl.NewCategoryUsecase(categoryRepository)
	categoryHandler := adminHandler.NewCategoryHandler(categoryUsecase)
	categoryDelivery := admin_delivery.NewCategoryDelivery(categoryHandler)

	courseRepository := repository_impl.NewCourseRepository(entClient)
	courseUsecase := usecase_impl.NewCourseUsecase(courseRepository)
	adminCourseHandler := adminHandler.NewCourseHandler(courseUsecase)
	adminCourseDelivery := admin_delivery.NewCourseDelivery(adminCourseHandler)

	rootDelivery := delivery.NewRootDelivery(authDelivery, categoryDelivery, adminCourseDelivery, authMiddleware)
	rootDelivery.RegisterRouter(ginEngine)
}
