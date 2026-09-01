package admin

import (
	"lms-api/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type QuizDelivery struct {
	quizHandler *admin.QuizHandler
}

func NewQuizDelivery(quizHandler *admin.QuizHandler) *QuizDelivery {
	return &QuizDelivery{quizHandler: quizHandler}
}

func (d *QuizDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	adminGroup.POST("lessons/:lessonId/quizzes", d.quizHandler.Create)
	adminGroup.GET("lessons/:lessonId/quizzes", d.quizHandler.FindByLessonID)

	quizzes := adminGroup.Group("quizzes")
	{
		quizzes.GET(":id", d.quizHandler.FindByID)
		quizzes.PUT(":id", d.quizHandler.Update)
		quizzes.DELETE(":id", d.quizHandler.Delete)
	}
}
