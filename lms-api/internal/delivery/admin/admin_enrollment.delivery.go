package admin

import (
	"lms-api/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type EnrollmentDelivery struct {
	enrollmentHandler *admin.EnrollmentHandler
}

func NewEnrollmentDelivery(enrollmentHandler *admin.EnrollmentHandler) *EnrollmentDelivery {
	return &EnrollmentDelivery{enrollmentHandler: enrollmentHandler}
}

func (d *EnrollmentDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	e := adminGroup.Group("enrollments")
	{
		e.GET("", d.enrollmentHandler.FindAll)
		e.GET(":id", d.enrollmentHandler.FindByID)
		e.PATCH(":id/status", d.enrollmentHandler.UpdateStatus)
	}
}
