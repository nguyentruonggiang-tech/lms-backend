package repository

import (
	"context"
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
)

type CategoryRepository interface {
	Create(ctx context.Context, body dto.CategoryCreateReq) (*ent.Categories, error)
	FindAll(ctx context.Context, query pagination.Query) ([]*ent.Categories, error)
	Count(ctx context.Context) (int, error)
	FindByID(ctx context.Context, id int) (*ent.Categories, error)
	Update(ctx context.Context, id int, body dto.CategoryUpdateReq) (*ent.Categories, error)
	Delete(ctx context.Context, id int) error
}
