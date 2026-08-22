package repository

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/dto"
)

type CourseRepository interface {
	Create(ctx context.Context, body dto.CourseCreateReq) (*ent.Courses, error)
	FindAll(ctx context.Context, query pagination.Query) ([]*ent.Courses, error)
	Count(ctx context.Context) (int, error)
	FindByID(ctx context.Context, id int) (*ent.Courses, error)
	Update(ctx context.Context, id int, body dto.CourseUpdateReq) (*ent.Courses, error)
	UpdateStatus(ctx context.Context, id int, status string) (*ent.Courses, error)
	Delete(ctx context.Context, id int) error
}
