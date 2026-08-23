package delivery

import (
	"lms-backend/internal/common/middlewares"
	adminDelivery "lms-backend/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	authDelivery          *authDelivery
	categoryDelivery      *categoryDelivery
	publicCourseDelivery  *courseDelivery
	enrollmentDelivery    *enrollmentDelivery
	adminCategoryDelivery *adminDelivery.CategoryDelivery
	adminCourseDelivery   *adminDelivery.CourseDelivery
	sectionDelivery       *adminDelivery.SectionDelivery
	lessonDelivery        *adminDelivery.LessonDelivery
	userDelivery          *adminDelivery.UserDelivery
	quizDelivery          *adminDelivery.QuizDelivery
	questionDelivery      *adminDelivery.QuestionDelivery
	authMiddleware        *middlewares.AuthMiddleware
}

func NewRootDelivery(
	authDelivery *authDelivery,
	categoryDelivery *categoryDelivery,
	publicCourseDelivery *courseDelivery,
	enrollmentDelivery *enrollmentDelivery,
	adminCategoryDelivery *adminDelivery.CategoryDelivery,
	adminCourseDelivery *adminDelivery.CourseDelivery,
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
		publicCourseDelivery:  publicCourseDelivery,
		enrollmentDelivery:    enrollmentDelivery,
		adminCategoryDelivery: adminCategoryDelivery,
		adminCourseDelivery:   adminCourseDelivery,
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
		r.publicCourseDelivery.RegisterRouter(apiGroup)

		studentGroup := apiGroup.Group("")
		studentGroup.Use(r.authMiddleware.Protect)
		r.enrollmentDelivery.RegisterRouter(studentGroup)

		adminGroup := apiGroup.Group("admin")
		adminGroup.Use(r.authMiddleware.Protect, r.authMiddleware.AdminOnly)
		{
			r.userDelivery.RegisterRouter(adminGroup)
			r.adminCategoryDelivery.RegisterRouter(adminGroup)
			r.adminCourseDelivery.RegisterRouter(adminGroup)
			r.sectionDelivery.RegisterRouter(adminGroup)
			r.lessonDelivery.RegisterRouter(adminGroup)
			r.quizDelivery.RegisterRouter(adminGroup)
			r.questionDelivery.RegisterRouter(adminGroup)
		}
	}
}
