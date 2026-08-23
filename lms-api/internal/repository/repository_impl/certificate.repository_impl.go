package repository_impl

import (
	"context"

	"lms-api/ent"
	"lms-api/ent/certificates"
	"lms-api/internal/common/pagination"
	"lms-api/internal/repository"
)

type certificateRepository struct {
	client *ent.Client
}

func NewCertificateRepository(client *ent.Client) repository.CertificateRepository {
	return &certificateRepository{client: client}
}

func (r *certificateRepository) Create(ctx context.Context, userID, courseID int, code string) (*ent.Certificates, error) {
	return r.client.Certificates.Create().
		SetUserID(userID).
		SetCourseID(courseID).
		SetCode(code).
		Save(ctx)
}

func (r *certificateRepository) FindAllByUserID(ctx context.Context, userID int, query pagination.Query) ([]*ent.Certificates, error) {
	return r.client.Certificates.Query().
		WithCourses().
		Where(certificates.UserIDEQ(userID)).
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *certificateRepository) CountByUserID(ctx context.Context, userID int) (int, error) {
	return r.client.Certificates.Query().
		Where(certificates.UserIDEQ(userID)).
		Count(ctx)
}

func (r *certificateRepository) FindByUserAndCourse(ctx context.Context, userID, courseID int) (*ent.Certificates, error) {
	return r.client.Certificates.Query().
		WithCourses().
		Where(
			certificates.UserIDEQ(userID),
			certificates.CourseIDEQ(courseID),
		).
		Only(ctx)
}

func (r *certificateRepository) FindByCode(ctx context.Context, code string) (*ent.Certificates, error) {
	return r.client.Certificates.Query().
		WithCourses().
		WithUsers().
		Where(certificates.CodeEQ(code)).
		Only(ctx)
}

func (r *certificateRepository) ExistsByUserAndCourse(ctx context.Context, userID, courseID int) (bool, error) {
	return r.client.Certificates.Query().
		Where(
			certificates.UserIDEQ(userID),
			certificates.CourseIDEQ(courseID),
		).
		Exist(ctx)
}
