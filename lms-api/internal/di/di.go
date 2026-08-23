package di

import (
	"lms-api/ent"
	"lms-api/internal/common/cache"
	"lms-api/internal/common/elastic"
	"lms-api/internal/common/env"
	"lms-api/internal/common/middlewares"
	"lms-api/internal/common/rabbitmq"
	"lms-api/internal/delivery"
	adminDelivery "lms-api/internal/delivery/admin"
	"lms-api/internal/handler"
	adminHandler "lms-api/internal/handler/admin"
	"lms-api/internal/repository/repository_impl"
	"lms-api/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, e *env.Env) {
	tokenUsecase := usecase_impl.NewTokenUsecase(e)
	redisClient := cache.NewRedisClient(e)
	elasticClient := elastic.NewElasticClient(e)
	rabbitmqClient := rabbitmq.NewRabbitMQ(e)

	userRepository := repository_impl.NewUserRepository(entClient)
	authMiddleware := middlewares.NewAuthMiddleware(tokenUsecase, userRepository)

	authUsecase := usecase_impl.NewAuthUsecase(userRepository, tokenUsecase, redisClient, e)
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

	sectionRepository := repository_impl.NewSectionRepository(entClient)
	sectionUsecase := usecase_impl.NewSectionUsecase(sectionRepository, redisClient)
	adminSectionHandler := adminHandler.NewSectionHandler(sectionUsecase)
	adminSectionDelivery := adminDelivery.NewSectionDelivery(adminSectionHandler)

	lessonRepository := repository_impl.NewLessonRepository(entClient)
	lessonUsecase := usecase_impl.NewLessonUsecase(lessonRepository, sectionRepository, redisClient)
	adminLessonHandler := adminHandler.NewLessonHandler(lessonUsecase)
	adminLessonDelivery := adminDelivery.NewLessonDelivery(adminLessonHandler)

	courseRepository := repository_impl.NewCourseRepository(entClient)
	courseUsecase := usecase_impl.NewCourseUsecase(courseRepository, lessonRepository, redisClient, elasticClient)
	publicCourseHandler := handler.NewCourseHandler(courseUsecase)
	publicCourseDelivery := delivery.NewCourseDelivery(publicCourseHandler)
	adminCourseHandler := adminHandler.NewCourseHandler(courseUsecase)
	adminCourseDelivery := adminDelivery.NewCourseDelivery(adminCourseHandler)

	quizRepository := repository_impl.NewQuizRepository(entClient)
	quizUsecase := usecase_impl.NewQuizUsecase(quizRepository, lessonRepository)
	adminQuizHandler := adminHandler.NewQuizHandler(quizUsecase)
	adminQuizDelivery := adminDelivery.NewQuizDelivery(adminQuizHandler)

	questionRepository := repository_impl.NewQuestionRepository(entClient)
	questionUsecase := usecase_impl.NewQuestionUsecase(questionRepository)
	adminQuestionHandler := adminHandler.NewQuestionHandler(questionUsecase)
	adminQuestionDelivery := adminDelivery.NewQuestionDelivery(adminQuestionHandler)

	enrollmentRepository := repository_impl.NewEnrollmentRepository(entClient)
	enrollmentUsecase := usecase_impl.NewEnrollmentUsecase(enrollmentRepository, courseRepository, rabbitmqClient)
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentUsecase)
	enrollmentDelivery := delivery.NewEnrollmentDelivery(enrollmentHandler)

	lessonProgressRepository := repository_impl.NewLessonProgressRepository(entClient)
	lessonProgressUsecase := usecase_impl.NewLessonProgressUsecase(lessonProgressRepository, enrollmentRepository, lessonRepository)
	lessonProgressHandler := handler.NewLessonProgressHandler(lessonProgressUsecase)
	lessonProgressDelivery := delivery.NewLessonProgressDelivery(lessonProgressHandler)

	adminEnrollmentUsecase := usecase_impl.NewAdminEnrollmentUsecase(enrollmentRepository)
	adminEnrollmentHandler := adminHandler.NewEnrollmentHandler(adminEnrollmentUsecase)
	adminEnrollmentDelivery := adminDelivery.NewEnrollmentDelivery(adminEnrollmentHandler)

	quizAttemptRepository := repository_impl.NewQuizAttemptRepository(entClient)
	quizClientUsecase := usecase_impl.NewQuizClientUsecase(quizRepository, questionRepository, quizAttemptRepository, enrollmentRepository)
	quizClientHandler := handler.NewQuizClientHandler(quizClientUsecase)
	quizClientDelivery := delivery.NewQuizClientDelivery(quizClientHandler)

	rootDelivery := delivery.NewRootDelivery(authDelivery, categoryDelivery, publicCourseDelivery, enrollmentDelivery, lessonProgressDelivery, quizClientDelivery, adminCategoryDelivery, adminCourseDelivery, adminSectionDelivery, adminLessonDelivery, adminUserDelivery, adminQuizDelivery, adminQuestionDelivery, adminEnrollmentDelivery, authMiddleware)
	rootDelivery.RegisterRouter(ginEngine)
}
