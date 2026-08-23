package delivery

import (
	"lms-backend/internal/common/middlewares"
	adminDelivery "lms-backend/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	authDelivery         *authDelivery
	categoryDelivery     *categoryDelivery
	adminCategoryDelivery *adminDelivery.CategoryDelivery
	courseDelivery       *adminDelivery.CourseDelivery
	sectionDelivery      *adminDelivery.SectionDelivery
	lessonDelivery       *adminDelivery.LessonDelivery
	userDelivery         *adminDelivery.UserDelivery
	quizDelivery         *adminDelivery.QuizDelivery
	questionDelivery     *adminDelivery.QuestionDelivery
	authMiddleware       *middlewares.AuthMiddleware
}

func NewRootDelivery(
	authDelivery *authDelivery,
	categoryDelivery *categoryDelivery,
	adminCategoryDelivery *adminDelivery.CategoryDelivery,
	courseDelivery *adminDelivery.CourseDelivery,
	sectionDelivery *adminDelivery.SectionDelivery,
	lessonDelivery *adminDelivery.LessonDelivery,
	userDelivery *adminDelivery.UserDelivery,
	quizDelivery *adminDelivery.QuizDelivery,
	questionDelivery *adminDelivery.QuestionDelivery,
	authMiddleware *middlewares.AuthMiddleware,
) *rootDelivery {
	return &rootDelivery{
		authDelivery:          authDelivery,
		categoryDelivery:      categoryDelivery,
		adminCategoryDelivery: adminCategoryDelivery,
		courseDelivery:        courseDelivery,
		sectionDelivery:       sectionDelivery,
		lessonDelivery:        lessonDelivery,
		userDelivery:          userDelivery,
		quizDelivery:          quizDelivery,
		questionDelivery:      questionDelivery,
		authMiddleware:        authMiddleware,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{
		r.authDelivery.RegisterRouter(apiGroup)
		r.categoryDelivery.RegisterRouter(apiGroup)

		adminGroup := apiGroup.Group("admin")
		adminGroup.Use(r.authMiddleware.Protect, r.authMiddleware.AdminOnly)
		{
			r.userDelivery.RegisterRouter(adminGroup)
			r.adminCategoryDelivery.RegisterRouter(adminGroup)
			r.courseDelivery.RegisterRouter(adminGroup)
			r.sectionDelivery.RegisterRouter(adminGroup)
			r.lessonDelivery.RegisterRouter(adminGroup)
			r.quizDelivery.RegisterRouter(adminGroup)
			r.questionDelivery.RegisterRouter(adminGroup)
		}
	}
}
