package repository

import (
	"context"
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
)

type QuizRepository interface {
	Create(ctx context.Context, lessonID, courseID int, body dto.QuizCreateReq) (*ent.Quizzes, error)
	FindByLessonID(ctx context.Context, lessonID int, query pagination.Query) ([]*ent.Quizzes, error)
	CountByLessonID(ctx context.Context, lessonID int) (int, error)
	FindByCourseID(ctx context.Context, courseID int, query pagination.Query) ([]*ent.Quizzes, error)
	CountByCourseID(ctx context.Context, courseID int) (int, error)
	FindByID(ctx context.Context, id int) (*ent.Quizzes, error)
	Update(ctx context.Context, id int, body dto.QuizUpdateReq) (*ent.Quizzes, error)
	Delete(ctx context.Context, id int) error
}
