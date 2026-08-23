package delivery

import (
	"lms-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type lessonProgressDelivery struct {
	lessonProgressHandler *handler.LessonProgressHandler
}

func NewLessonProgressDelivery(lessonProgressHandler *handler.LessonProgressHandler) *lessonProgressDelivery {
	return &lessonProgressDelivery{lessonProgressHandler: lessonProgressHandler}
}

func (d *lessonProgressDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	myGroup := apiGroup.Group("my")
	{
		myGroup.GET("courses/:courseId/lessons", d.lessonProgressHandler.ListLessons)
		myGroup.GET("lessons/:lessonId", d.lessonProgressHandler.FindLesson)
		myGroup.POST("lessons/:lessonId/complete", d.lessonProgressHandler.Complete)
		myGroup.GET("courses/:courseId/progress", d.lessonProgressHandler.GetProgress)
	}
}
