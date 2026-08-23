package usecase

import (
	"context"
	"lms-backend/internal/dto"
)

type LessonUsecase interface {
	Create(ctx context.Context, sectionID int, body dto.LessonCreateReq) (any, error)
	FindByCourseID(ctx context.Context, courseID int, page, limit string) (any, error)
	FindByID(ctx context.Context, id int) (any, error)
	Update(ctx context.Context, id int, body dto.LessonUpdateReq) (any, error)
	Delete(ctx context.Context, id int) (any, error)
}
