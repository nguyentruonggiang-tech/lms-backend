package repository

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, userID, courseID int) (*ent.Enrollments, error)
	ExistsByUserAndCourse(ctx context.Context, userID, courseID int) (bool, error)
	FindAllByUserID(ctx context.Context, userID int, status string, query pagination.Query) ([]*ent.Enrollments, error)
	CountByUserID(ctx context.Context, userID int, status string) (int, error)
	FindByUserAndCourse(ctx context.Context, userID, courseID int) (*ent.Enrollments, error)
	FindByID(ctx context.Context, id int) (*ent.Enrollments, error)
	Delete(ctx context.Context, id int) error
}
