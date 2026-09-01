package admin

import (
	"lms-api/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type QuestionDelivery struct {
	questionHandler *admin.QuestionHandler
}

func NewQuestionDelivery(questionHandler *admin.QuestionHandler) *QuestionDelivery {
	return &QuestionDelivery{questionHandler: questionHandler}
}

func (d *QuestionDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	quizQuestions := adminGroup.Group("quizzes/:quizId/questions")
	{
		quizQuestions.POST("", d.questionHandler.Create)
		quizQuestions.GET("", d.questionHandler.FindByQuizID)
	}

	questions := adminGroup.Group("questions")
	{
		questions.PUT(":id", d.questionHandler.Update)
		questions.DELETE(":id", d.questionHandler.Delete)
	}
}
