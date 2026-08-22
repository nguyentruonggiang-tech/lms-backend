package usecase

import (
	"context"
	"lms-backend/internal/dto"
)

type CategoryUsecase interface {
	Create(ctx context.Context, body dto.CategoryCreateReq) (any, error)
	FindAll(ctx context.Context, page, limit string) (any, error)
	FindByID(ctx context.Context, id int) (any, error)
	Update(ctx context.Context, id int, body dto.CategoryUpdateReq) (any, error)
	Delete(ctx context.Context, id int) (any, error)
}
