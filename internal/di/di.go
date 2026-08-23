package di

import (
	"lms-backend/ent"
	"lms-backend/internal/common/env"
	"lms-backend/internal/common/middlewares"
	"lms-backend/internal/delivery"
	adminDelivery "lms-backend/internal/delivery/admin"
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

	adminUserUsecase := usecase_impl.NewAdminUserUsecase(userRepository)
	adminUserHandler := adminHandler.NewUserHandler(adminUserUsecase)
	adminUserDelivery := adminDelivery.NewUserDelivery(adminUserHandler)

	categoryRepository := repository_impl.NewCategoryRepository(entClient)
	categoryUsecase := usecase_impl.NewCategoryUsecase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	categoryDelivery := delivery.NewCategoryDelivery(categoryHandler)
	adminCategoryHandler := adminHandler.NewCategoryHandler(categoryUsecase)
	adminCategoryDelivery := adminDelivery.NewCategoryDelivery(adminCategoryHandler)

	courseRepository := repository_impl.NewCourseRepository(entClient)
	courseUsecase := usecase_impl.NewCourseUsecase(courseRepository)
	adminCourseHandler := adminHandler.NewCourseHandler(courseUsecase)
	adminCourseDelivery := adminDelivery.NewCourseDelivery(adminCourseHandler)

	sectionRepository := repository_impl.NewSectionRepository(entClient)
	sectionUsecase := usecase_impl.NewSectionUsecase(sectionRepository)
	adminSectionHandler := adminHandler.NewSectionHandler(sectionUsecase)
	adminSectionDelivery := adminDelivery.NewSectionDelivery(adminSectionHandler)

	lessonRepository := repository_impl.NewLessonRepository(entClient)
	lessonUsecase := usecase_impl.NewLessonUsecase(lessonRepository, sectionRepository)
	adminLessonHandler := adminHandler.NewLessonHandler(lessonUsecase)
	adminLessonDelivery := adminDelivery.NewLessonDelivery(adminLessonHandler)

	quizRepository := repository_impl.NewQuizRepository(entClient)
	quizUsecase := usecase_impl.NewQuizUsecase(quizRepository, lessonRepository)
	adminQuizHandler := adminHandler.NewQuizHandler(quizUsecase)
	adminQuizDelivery := adminDelivery.NewQuizDelivery(adminQuizHandler)

	questionRepository := repository_impl.NewQuestionRepository(entClient)
	questionUsecase := usecase_impl.NewQuestionUsecase(questionRepository)
	adminQuestionHandler := adminHandler.NewQuestionHandler(questionUsecase)
	adminQuestionDelivery := adminDelivery.NewQuestionDelivery(adminQuestionHandler)

	rootDelivery := delivery.NewRootDelivery(authDelivery, categoryDelivery, adminCategoryDelivery, adminCourseDelivery, adminSectionDelivery, adminLessonDelivery, adminUserDelivery, adminQuizDelivery, adminQuestionDelivery, authMiddleware)
	rootDelivery.RegisterRouter(ginEngine)
}
