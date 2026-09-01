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
	adminGroup.GET("courses/:courseId/quizzes", d.quizHandler.FindByCourseID)

	quizzes := adminGroup.Group("quizzes")
	{
		quizzes.GET(":quizId", d.quizHandler.FindByID)
		quizzes.PUT(":quizId", d.quizHandler.Update)
		quizzes.DELETE(":quizId", d.quizHandler.Delete)
	}
}
