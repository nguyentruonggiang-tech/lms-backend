package usecase

import (
	"context"
	"lms-backend/internal/dto"
)

type QuizUsecase interface {
	Create(ctx context.Context, lessonID int, body dto.QuizCreateReq) (any, error)
	FindByLessonID(ctx context.Context, lessonID int, page, limit string) (any, error)
	FindByID(ctx context.Context, id int) (any, error)
	Update(ctx context.Context, id int, body dto.QuizUpdateReq) (any, error)
	Delete(ctx context.Context, id int) (any, error)
}
