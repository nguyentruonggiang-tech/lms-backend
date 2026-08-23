package repository

import (
	"context"
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, userID, courseID int) (*ent.Enrollments, error)
	ExistsByUserAndCourse(ctx context.Context, userID, courseID int) (bool, error)
	FindAllByUserID(ctx context.Context, userID int, status string, query pagination.Query) ([]*ent.Enrollments, error)
	CountByUserID(ctx context.Context, userID int, status string) (int, error)
	FindByUserAndCourse(ctx context.Context, userID, courseID int) (*ent.Enrollments, error)
	FindByID(ctx context.Context, id int) (*ent.Enrollments, error)
	Delete(ctx context.Context, id int) error
	UpdateProgressPercent(ctx context.Context, userID, courseID int, percent float64) error
	CompleteEnrollment(ctx context.Context, userID, courseID int) error

	FindAll(ctx context.Context, filter dto.AdminEnrollmentFilter, query pagination.Query) ([]*ent.Enrollments, error)
	CountAll(ctx context.Context, filter dto.AdminEnrollmentFilter) (int, error)
	UpdateStatus(ctx context.Context, id int, status string) (*ent.Enrollments, error)
}
