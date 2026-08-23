package usecase

import (
	"context"
	"lms-api/internal/dto"
)

type QuizClientUsecase interface {
	GetQuiz(ctx context.Context, userID, quizID int) (any, error)
	Submit(ctx context.Context, userID, quizID int, body dto.QuizSubmitReq) (any, error)
	GetAttempts(ctx context.Context, userID, quizID int, page, limit string) (any, error)
}
