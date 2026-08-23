package usecase

import (
	"context"
	"lms-api/internal/dto"
)

type QuestionUsecase interface {
	Create(ctx context.Context, quizID int, body dto.QuestionCreateReq) (any, error)
	FindByQuizID(ctx context.Context, quizID int, page, limit string) (any, error)
	Update(ctx context.Context, id int, body dto.QuestionUpdateReq) (any, error)
	Delete(ctx context.Context, id int) (any, error)
}
