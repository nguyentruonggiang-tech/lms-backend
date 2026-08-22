package repository_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/ent/courses"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
)

type courseRepository struct {
	client *ent.Client
}

func NewCourseRepository(client *ent.Client) repository.CourseRepository {
	return &courseRepository{client: client}
}

func (r *courseRepository) Create(ctx context.Context, body dto.CourseCreateReq) (*ent.Courses, error) {
	q := r.client.Courses.Create().
		SetCategoryID(body.CategoryID).
		SetTitle(body.Title).
		SetSlug(body.Slug).
		SetPrice(body.Price)

	if body.Description != nil {
		q = q.SetDescription(*body.Description)
	}
	if body.Thumbnail != nil {
		q = q.SetThumbnail(*body.Thumbnail)
	}
	if body.Level != "" {
		q = q.SetLevel(courses.Level(body.Level))
	}
	if body.Status != "" {
		q = q.SetStatus(courses.Status(body.Status))
	}

	return q.Save(ctx)
}

func (r *courseRepository) FindAll(ctx context.Context, query pagination.Query) ([]*ent.Courses, error) {
	return r.client.Courses.Query().
		WithCategories().
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *courseRepository) Count(ctx context.Context) (int, error) {
	return r.client.Courses.Query().Count(ctx)
}

func (r *courseRepository) FindByID(ctx context.Context, id int) (*ent.Courses, error) {
	return r.client.Courses.Query().
		Where(courses.IDEQ(id)).
		WithCategories().
		Only(ctx)
}

func (r *courseRepository) Update(ctx context.Context, id int, body dto.CourseUpdateReq) (*ent.Courses, error) {
	q := r.client.Courses.UpdateOneID(id)

	if body.CategoryID != nil {
		q = q.SetCategoryID(*body.CategoryID)
	}
	if body.Title != nil {
		q = q.SetTitle(*body.Title)
	}
	if body.Slug != nil {
		q = q.SetSlug(*body.Slug)
	}
	if body.Description != nil {
		q = q.SetDescription(*body.Description)
	}
	if body.Thumbnail != nil {
		q = q.SetThumbnail(*body.Thumbnail)
	}
	if body.Price != nil {
		q = q.SetPrice(*body.Price)
	}
	if body.Level != nil {
		q = q.SetLevel(courses.Level(*body.Level))
	}
	if body.Status != nil {
		q = q.SetStatus(courses.Status(*body.Status))
	}

	return q.Save(ctx)
}

func (r *courseRepository) UpdateStatus(ctx context.Context, id int, status string) (*ent.Courses, error) {
	return r.client.Courses.UpdateOneID(id).
		SetStatus(courses.Status(status)).
		Save(ctx)
}

func (r *courseRepository) Delete(ctx context.Context, id int) error {
	return r.client.Courses.DeleteOneID(id).Exec(ctx)
}
