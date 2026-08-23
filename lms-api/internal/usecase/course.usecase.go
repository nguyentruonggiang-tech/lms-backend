package usecase

import (
	"context"
	"lms-api/internal/dto"
)

type CourseUsecase interface {
	Create(ctx context.Context, body dto.CourseCreateReq) (any, error)
	FindAll(ctx context.Context, page, limit string) (any, error)
	FindByID(ctx context.Context, id int) (any, error)
	Update(ctx context.Context, id int, body dto.CourseUpdateReq) (any, error)
	UpdateStatus(ctx context.Context, id int, body dto.CourseUpdateStatusReq) (any, error)
	Delete(ctx context.Context, id int) (any, error)

	FindAllPublished(ctx context.Context, filter dto.CoursePublicFilter, page, limit string) (any, error)
	SearchPublished(ctx context.Context, filter dto.CoursePublicFilter, page, limit string) (any, error)
	FindPublishedByID(ctx context.Context, id int) (any, error)
	FindPreviewLessons(ctx context.Context, courseID int) (any, error)
}
