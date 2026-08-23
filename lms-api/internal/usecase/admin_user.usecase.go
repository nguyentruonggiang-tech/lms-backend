package usecase

import (
	"context"
	"lms-api/internal/dto"
)

type AdminUserUsecase interface {
	FindAll(ctx context.Context, page, limit, role, status, keyword string) (any, error)
	FindByID(ctx context.Context, id int) (any, error)
	UpdateStatus(ctx context.Context, id int, body dto.UserUpdateStatusReq) (any, error)
	UpdateRole(ctx context.Context, id int, body dto.UserUpdateRoleReq) (any, error)
}
