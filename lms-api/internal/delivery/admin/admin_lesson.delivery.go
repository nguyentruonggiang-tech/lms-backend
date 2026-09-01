package admin

import (
	"lms-api/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type LessonDelivery struct {
	lessonHandler *admin.LessonHandler
}

func NewLessonDelivery(lessonHandler *admin.LessonHandler) *LessonDelivery {
	return &LessonDelivery{lessonHandler: lessonHandler}
}

func (d *LessonDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	adminGroup.POST("sections/:sectionId/lessons", d.lessonHandler.Create)
	adminGroup.GET("courses/:courseId/lessons", d.lessonHandler.FindByCourseID)

	lessons := adminGroup.Group("lessons")
	{
		lessons.GET(":id", d.lessonHandler.FindByID)
		lessons.PUT(":id", d.lessonHandler.Update)
		lessons.DELETE(":id", d.lessonHandler.Delete)
	}
}
