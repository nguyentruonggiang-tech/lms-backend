package repository

import (
	"context"

	"lms-api/ent"
	"lms-api/internal/common/pagination"
)

type CertificateRepository interface {
	Create(ctx context.Context, userID, courseID int, code string) (*ent.Certificates, error)
	FindAllByUserID(ctx context.Context, userID int, query pagination.Query) ([]*ent.Certificates, error)
	CountByUserID(ctx context.Context, userID int) (int, error)
	FindByUserAndCourse(ctx context.Context, userID, courseID int) (*ent.Certificates, error)
	FindByCode(ctx context.Context, code string) (*ent.Certificates, error)
	ExistsByUserAndCourse(ctx context.Context, userID, courseID int) (bool, error)
}
