package repository

import (
	"context"
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
)

type CourseRepository interface {
	Create(ctx context.Context, body dto.CourseCreateReq) (*ent.Courses, error)
	FindAll(ctx context.Context, query pagination.Query) ([]*ent.Courses, error)
	Count(ctx context.Context) (int, error)
	FindByID(ctx context.Context, id int) (*ent.Courses, error)
	Update(ctx context.Context, id int, body dto.CourseUpdateReq) (*ent.Courses, error)
	UpdateStatus(ctx context.Context, id int, status string) (*ent.Courses, error)
	Delete(ctx context.Context, id int) error

	FindAllPublished(ctx context.Context, filter dto.CoursePublicFilter, query pagination.Query) ([]*ent.Courses, error)
	CountPublished(ctx context.Context, filter dto.CoursePublicFilter) (int, error)
	SearchPublished(ctx context.Context, filter dto.CoursePublicFilter, query pagination.Query) ([]*ent.Courses, error)
	CountSearch(ctx context.Context, filter dto.CoursePublicFilter) (int, error)
	FindPublishedByID(ctx context.Context, id int) (*ent.Courses, error)
}
