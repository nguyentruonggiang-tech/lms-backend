package repository

import (
	"context"
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
)

type SectionRepository interface {
	Create(ctx context.Context, courseID int, body dto.SectionCreateReq) (*ent.Sections, error)
	FindByID(ctx context.Context, id int) (*ent.Sections, error)
	FindByCourseID(ctx context.Context, courseID int, query pagination.Query) ([]*ent.Sections, error)
	CountByCourseID(ctx context.Context, courseID int) (int, error)
	Update(ctx context.Context, id int, body dto.SectionUpdateReq) (*ent.Sections, error)
	Delete(ctx context.Context, id int) error
}
