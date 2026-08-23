package admin

import (
	"lms-backend/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type CourseDelivery struct {
	courseHandler *admin.CourseHandler
}

func NewCourseDelivery(courseHandler *admin.CourseHandler) *CourseDelivery {
	return &CourseDelivery{
		courseHandler: courseHandler,
	}
}

func (d *CourseDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	courses := adminGroup.Group("courses")
	{
		courses.POST("", d.courseHandler.Create)
		courses.GET("", d.courseHandler.FindAll)
		courses.GET(":id", d.courseHandler.FindByID)
		courses.PUT(":id", d.courseHandler.Update)
		courses.PATCH(":id/status", d.courseHandler.UpdateStatus)
		courses.DELETE(":id", d.courseHandler.Delete)
	}
}
