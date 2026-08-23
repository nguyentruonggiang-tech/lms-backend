package repository

import (
	"context"
	"lms-api/internal/dto"
)

type DashboardRepository interface {
	GetOverview(ctx context.Context) (dto.DashboardOverview, error)
	GetTopCourses(ctx context.Context, filter dto.TopCourseFilter) ([]dto.TopCourseItem, error)
}
