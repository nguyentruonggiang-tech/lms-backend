package usecase

import (
	"context"
	"lms-api/internal/dto"
)

type DashboardUsecase interface {
	GetOverview(ctx context.Context) (any, error)
	GetTopCourses(ctx context.Context, req dto.TopCourseReq) (any, error)
}
