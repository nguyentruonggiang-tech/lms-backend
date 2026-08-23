package admin_delivery

import (
	"lms-backend/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type LessonDelivery struct {
	lessonHandler *admin.LessonHandler
}

func NewLessonDelivery(lessonHandler *admin.LessonHandler) *LessonDelivery {
	return &LessonDelivery{lessonHandler: lessonHandler}
}

func (d *LessonDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	adminGroup.POST("sections/:section_id/lessons", d.lessonHandler.Create)
	adminGroup.GET("courses/:course_id/lessons", d.lessonHandler.FindByCourseID)

	lessons := adminGroup.Group("lessons")
	{
		lessons.GET(":id", d.lessonHandler.FindByID)
		lessons.PUT(":id", d.lessonHandler.Update)
		lessons.DELETE(":id", d.lessonHandler.Delete)
	}
}
