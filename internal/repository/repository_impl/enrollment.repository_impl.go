package repository_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/ent/enrollments"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/repository"
)

type enrollmentRepository struct {
	client *ent.Client
}

func NewEnrollmentRepository(client *ent.Client) repository.EnrollmentRepository {
	return &enrollmentRepository{client: client}
}

func (r *enrollmentRepository) Create(ctx context.Context, userID, courseID int) (*ent.Enrollments, error) {
	return r.client.Enrollments.Create().
		SetUserID(userID).
		SetCourseID(courseID).
		Save(ctx)
}

func (r *enrollmentRepository) ExistsByUserAndCourse(ctx context.Context, userID, courseID int) (bool, error) {
	return r.client.Enrollments.Query().
		Where(
			enrollments.UserIDEQ(userID),
			enrollments.CourseIDEQ(courseID),
		).
		Exist(ctx)
}

func (r *enrollmentRepository) FindAllByUserID(ctx context.Context, userID int, status string, query pagination.Query) ([]*ent.Enrollments, error) {
	q := r.client.Enrollments.Query().
		WithCourses().
		Where(enrollments.UserIDEQ(userID))

	if status != "" {
		q = q.Where(enrollments.StatusEQ(enrollments.Status(status)))
	}

	return q.Limit(query.Limit).Offset(query.Offset).All(ctx)
}

func (r *enrollmentRepository) CountByUserID(ctx context.Context, userID int, status string) (int, error) {
	q := r.client.Enrollments.Query().
		Where(enrollments.UserIDEQ(userID))

	if status != "" {
		q = q.Where(enrollments.StatusEQ(enrollments.Status(status)))
	}

	return q.Count(ctx)
}

func (r *enrollmentRepository) FindByUserAndCourse(ctx context.Context, userID, courseID int) (*ent.Enrollments, error) {
	return r.client.Enrollments.Query().
		WithCourses().
		Where(
			enrollments.UserIDEQ(userID),
			enrollments.CourseIDEQ(courseID),
		).
		Only(ctx)
}

func (r *enrollmentRepository) FindByID(ctx context.Context, id int) (*ent.Enrollments, error) {
	return r.client.Enrollments.Query().
		Where(enrollments.IDEQ(id)).
		Only(ctx)
}

func (r *enrollmentRepository) Delete(ctx context.Context, id int) error {
	return r.client.Enrollments.DeleteOneID(id).Exec(ctx)
}

func (r *enrollmentRepository) UpdateProgressPercent(ctx context.Context, userID, courseID int, percent float64) error {
	return r.client.Enrollments.Update().
		Where(
			enrollments.UserIDEQ(userID),
			enrollments.CourseIDEQ(courseID),
		).
		SetProgressPercent(percent).
		Exec(ctx)
}
