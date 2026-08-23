package usecase

import (
	"context"
	"lms-backend/internal/dto"
)

type SectionUsecase interface {
	Create(ctx context.Context, courseID int, body dto.SectionCreateReq) (any, error)
	FindByCourseID(ctx context.Context, courseID int, page, limit string) (any, error)
	Update(ctx context.Context, id int, body dto.SectionUpdateReq) (any, error)
	Delete(ctx context.Context, id int) (any, error)
}
