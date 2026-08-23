package usecase

import (
	"context"
	"lms-backend/internal/dto"
)

type AdminEnrollmentUsecase interface {
	FindAll(ctx context.Context, filter dto.AdminEnrollmentFilter, page, limit string) (any, error)
	FindByID(ctx context.Context, id int) (any, error)
	UpdateStatus(ctx context.Context, id int, body dto.EnrollmentUpdateStatusReq) (any, error)
}
