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
	quizDelivery     *admin_delivery.QuizDelivery
	authMiddleware   *middlewares.AuthMiddleware
}

func NewRootDelivery(
	authDelivery *authDelivery,
	categoryDelivery *admin_delivery.CategoryDelivery,
	courseDelivery *admin_delivery.CourseDelivery,
	sectionDelivery *admin_delivery.SectionDelivery,
	lessonDelivery *admin_delivery.LessonDelivery,
	quizDelivery *admin_delivery.QuizDelivery,
	authMiddleware *middlewares.AuthMiddleware,
) *rootDelivery {
	return &rootDelivery{
		authDelivery:     authDelivery,
		categoryDelivery: categoryDelivery,
		courseDelivery:   courseDelivery,
		sectionDelivery:  sectionDelivery,
		lessonDelivery:   lessonDelivery,
		quizDelivery:     quizDelivery,
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
			r.categoryDelivery.RegisterRouter(adminGroup)
			r.courseDelivery.RegisterRouter(adminGroup)
			r.sectionDelivery.RegisterRouter(adminGroup)
			r.lessonDelivery.RegisterRouter(adminGroup)
			r.quizDelivery.RegisterRouter(adminGroup)
		}
	}
}
