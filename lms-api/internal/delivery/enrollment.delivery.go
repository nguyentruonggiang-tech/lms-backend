package delivery

import (
	"lms-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type enrollmentDelivery struct {
	enrollmentHandler *handler.EnrollmentHandler
}

func NewEnrollmentDelivery(enrollmentHandler *handler.EnrollmentHandler) *enrollmentDelivery {
	return &enrollmentDelivery{enrollmentHandler: enrollmentHandler}
}

func (d *enrollmentDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	apiGroup.POST("enrollments", d.enrollmentHandler.Enroll)

	myGroup := apiGroup.Group("my")
	{
		myGroup.GET("enrollments", d.enrollmentHandler.FindMyEnrollments)
		myGroup.GET("courses/:courseId", d.enrollmentHandler.FindMyCourse)
		myGroup.DELETE("enrollments/:id", d.enrollmentHandler.Cancel)
	}
}
