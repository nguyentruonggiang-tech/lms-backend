package usecase_impl

import (
	"context"
	"strconv"
	"time"

	"lms-api/internal/dto"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type dashboardUsecaseImpl struct {
	dashboardRepository repository.DashboardRepository
}

func NewDashboardUsecase(dashboardRepository repository.DashboardRepository) usecase.DashboardUsecase {
	return &dashboardUsecaseImpl{dashboardRepository: dashboardRepository}
}

func (u *dashboardUsecaseImpl) GetOverview(ctx context.Context) (any, error) {
	return u.dashboardRepository.GetOverview(ctx)
}

func (u *dashboardUsecaseImpl) GetTopCourses(ctx context.Context, req dto.TopCourseReq) (any, error) {
	filter := dto.TopCourseFilter{Limit: 10}

	if req.Limit != "" {
		if l, err := strconv.Atoi(req.Limit); err == nil && l > 0 {
			filter.Limit = l
		}
	}
	if req.FromDate != "" {
		if t, err := time.Parse("2006-01-02", req.FromDate); err == nil {
			filter.FromDate = &t
		}
	}
	if req.ToDate != "" {
		if t, err := time.Parse("2006-01-02", req.ToDate); err == nil {
			end := t.Add(24*time.Hour - time.Second)
			filter.ToDate = &end
		}
	}

	return u.dashboardRepository.GetTopCourses(ctx, filter)
}
