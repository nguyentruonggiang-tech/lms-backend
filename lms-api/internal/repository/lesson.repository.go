package repository

import (
	"context"
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
)

type LessonRepository interface {
	Create(ctx context.Context, sectionID, courseID int, body dto.LessonCreateReq) (*ent.Lessons, error)
	FindByCourseID(ctx context.Context, courseID int, query pagination.Query) ([]*ent.Lessons, error)
	CountByCourseID(ctx context.Context, courseID int) (int, error)
	FindByID(ctx context.Context, id int) (*ent.Lessons, error)
	Update(ctx context.Context, id int, body dto.LessonUpdateReq) (*ent.Lessons, error)
	Delete(ctx context.Context, id int) error
	FindPreviewByCourseID(ctx context.Context, courseID int) ([]*ent.Lessons, error)
}
