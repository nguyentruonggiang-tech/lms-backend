package delivery

import (
	"lms-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type quizClientDelivery struct {
	quizClientHandler *handler.QuizClientHandler
}

func NewQuizClientDelivery(quizClientHandler *handler.QuizClientHandler) *quizClientDelivery {
	return &quizClientDelivery{quizClientHandler: quizClientHandler}
}

func (d *quizClientDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	myGroup := apiGroup.Group("my")
	{
		myGroup.GET("quizzes/:quizId", d.quizClientHandler.GetQuiz)
		myGroup.POST("quizzes/:quizId/submit", d.quizClientHandler.Submit)
		myGroup.GET("quizzes/:quizId/attempts", d.quizClientHandler.GetAttempts)
	}
}
