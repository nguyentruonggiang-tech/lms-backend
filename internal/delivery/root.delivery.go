package delivery

import (
	"lms-backend/internal/common/middlewares"
	"lms-backend/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	authDelivery     *authDelivery
	categoryDelivery *admin_delivery.CategoryDelivery
	courseDelivery   *admin_delivery.CourseDelivery
	sectionDelivery  *admin_delivery.SectionDelivery
	lessonDelivery   *admin_delivery.LessonDelivery
	userDelivery     *admin_delivery.UserDelivery
	quizDelivery     *admin_delivery.QuizDelivery
	questionDelivery *admin_delivery.QuestionDelivery
	authMiddleware   *middlewares.AuthMiddleware
}

func NewRootDelivery(
	authDelivery *authDelivery,
	categoryDelivery *admin_delivery.CategoryDelivery,
	courseDelivery *admin_delivery.CourseDelivery,
	sectionDelivery *admin_delivery.SectionDelivery,
	lessonDelivery *admin_delivery.LessonDelivery,
	userDelivery *admin_delivery.UserDelivery,
	quizDelivery *admin_delivery.QuizDelivery,
	questionDelivery *admin_delivery.QuestionDelivery,
	authMiddleware *middlewares.AuthMiddleware,
) *rootDelivery {
	return &rootDelivery{
		authDelivery:     authDelivery,
		categoryDelivery: categoryDelivery,
		courseDelivery:   courseDelivery,
		sectionDelivery:  sectionDelivery,
		lessonDelivery:   lessonDelivery,
		userDelivery:     userDelivery,
		quizDelivery:     quizDelivery,
		questionDelivery: questionDelivery,
		authMiddleware:   authMiddleware,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{
		r.authDelivery.RegisterRouter(apiGroup)

		adminGroup := apiGroup.Group("admin")
		adminGroup.Use(r.authMiddleware.Protect, r.authMiddleware.AdminOnly)
		{
			r.userDelivery.RegisterRouter(adminGroup)
			r.categoryDelivery.RegisterRouter(adminGroup)
			r.courseDelivery.RegisterRouter(adminGroup)
			r.sectionDelivery.RegisterRouter(adminGroup)
			r.lessonDelivery.RegisterRouter(adminGroup)
			r.quizDelivery.RegisterRouter(adminGroup)
			r.questionDelivery.RegisterRouter(adminGroup)
		}
	}
}
