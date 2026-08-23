package repository

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/dto"
)

type QuestionRepository interface {
	Create(ctx context.Context, quizID int, body dto.QuestionCreateReq) (*ent.Questions, error)
	FindByQuizID(ctx context.Context, quizID int, query pagination.Query) ([]*ent.Questions, error)
	CountByQuizID(ctx context.Context, quizID int) (int, error)
	FindAllByQuizID(ctx context.Context, quizID int) ([]*ent.Questions, error)
	Update(ctx context.Context, id int, body dto.QuestionUpdateReq) (*ent.Questions, error)
	Delete(ctx context.Context, id int) error
}
