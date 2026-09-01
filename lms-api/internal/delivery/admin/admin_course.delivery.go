package admin

import (
	"lms-api/internal/handler/admin"

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
		courses.GET(":courseId", d.courseHandler.FindByID)
		courses.PUT(":courseId", d.courseHandler.Update)
		courses.PATCH(":courseId/status", d.courseHandler.UpdateStatus)
		courses.DELETE(":courseId", d.courseHandler.Delete)
		courses.POST(":courseId/reindex", d.courseHandler.Reindex)
	}
}
