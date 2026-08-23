package delivery

import (
	"lms-api/internal/common/middlewares"
	adminDelivery "lms-api/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	authDelivery            *authDelivery
	categoryDelivery        *categoryDelivery
	publicCourseDelivery    *courseDelivery
	enrollmentDelivery      *enrollmentDelivery
	lessonProgressDelivery  *lessonProgressDelivery
	quizClientDelivery      *quizClientDelivery
	certificateDelivery     *certificateDelivery
	adminCategoryDelivery   *adminDelivery.CategoryDelivery
	adminCourseDelivery     *adminDelivery.CourseDelivery
	sectionDelivery         *adminDelivery.SectionDelivery
	lessonDelivery          *adminDelivery.LessonDelivery
	userDelivery            *adminDelivery.UserDelivery
	quizDelivery            *adminDelivery.QuizDelivery
	questionDelivery        *adminDelivery.QuestionDelivery
	adminEnrollmentDelivery *adminDelivery.EnrollmentDelivery
	authMiddleware          *middlewares.AuthMiddleware
}

func NewRootDelivery(
	authDelivery *authDelivery,
	categoryDelivery *categoryDelivery,
	publicCourseDelivery *courseDelivery,
	enrollmentDelivery *enrollmentDelivery,
	lessonProgressDelivery *lessonProgressDelivery,
	quizClientDelivery *quizClientDelivery,
	certificateDelivery *certificateDelivery,
	adminCategoryDelivery *adminDelivery.CategoryDelivery,
	adminCourseDelivery *adminDelivery.CourseDelivery,
	sectionDelivery *adminDelivery.SectionDelivery,
	lessonDelivery *adminDelivery.LessonDelivery,
	userDelivery *adminDelivery.UserDelivery,
	quizDelivery *adminDelivery.QuizDelivery,
	questionDelivery *adminDelivery.QuestionDelivery,
	adminEnrollmentDelivery *adminDelivery.EnrollmentDelivery,
	authMiddleware *middlewares.AuthMiddleware,
) *rootDelivery {
	return &rootDelivery{
		authDelivery:            authDelivery,
		categoryDelivery:        categoryDelivery,
		publicCourseDelivery:    publicCourseDelivery,
		enrollmentDelivery:      enrollmentDelivery,
		lessonProgressDelivery:  lessonProgressDelivery,
		quizClientDelivery:      quizClientDelivery,
		certificateDelivery:     certificateDelivery,
		adminCategoryDelivery:   adminCategoryDelivery,
		adminCourseDelivery:     adminCourseDelivery,
		sectionDelivery:         sectionDelivery,
		lessonDelivery:          lessonDelivery,
		userDelivery:            userDelivery,
		quizDelivery:            quizDelivery,
		questionDelivery:        questionDelivery,
		adminEnrollmentDelivery: adminEnrollmentDelivery,
		authMiddleware:          authMiddleware,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{
		r.authDelivery.RegisterRouter(apiGroup)
		r.categoryDelivery.RegisterRouter(apiGroup)
		r.publicCourseDelivery.RegisterRouter(apiGroup)
		r.certificateDelivery.RegisterPublicRouter(apiGroup)

		studentGroup := apiGroup.Group("")
		studentGroup.Use(r.authMiddleware.Protect)
		r.enrollmentDelivery.RegisterRouter(studentGroup)
		r.lessonProgressDelivery.RegisterRouter(studentGroup)
		r.quizClientDelivery.RegisterRouter(studentGroup)
		r.certificateDelivery.RegisterStudentRouter(studentGroup)

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
			r.adminEnrollmentDelivery.RegisterRouter(adminGroup)
		}
	}
}
