package delivery

import (
	"lms-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type courseDelivery struct {
	courseHandler *handler.CourseHandler
}

func NewCourseDelivery(courseHandler *handler.CourseHandler) *courseDelivery {
	return &courseDelivery{courseHandler: courseHandler}
}

func (d *courseDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	courses := apiGroup.Group("courses")
	{
		courses.GET("", d.courseHandler.FindAllPublished)
		courses.GET("search", d.courseHandler.Search)
		courses.GET(":id", d.courseHandler.FindPublishedByID)
		courses.GET(":id/preview-lessons", d.courseHandler.FindPreviewLessons)
	}
}
