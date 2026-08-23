package repository_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/ent/sections"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
)

type sectionRepository struct {
	client *ent.Client
}

func NewSectionRepository(client *ent.Client) repository.SectionRepository {
	return &sectionRepository{client: client}
}

func (r *sectionRepository) Create(ctx context.Context, courseID int, body dto.SectionCreateReq) (*ent.Sections, error) {
	q := r.client.Sections.Create().
		SetCourseID(courseID).
		SetTitle(body.Title)

	if body.SortOrder != nil {
		q = q.SetSortOrder(*body.SortOrder)
	}

	return q.Save(ctx)
}

func (r *sectionRepository) FindByID(ctx context.Context, id int) (*ent.Sections, error) {
	return r.client.Sections.Get(ctx, id)
}

func (r *sectionRepository) FindByCourseID(ctx context.Context, courseID int, query pagination.Query) ([]*ent.Sections, error) {
	return r.client.Sections.Query().
		Where(sections.CourseIDEQ(courseID)).
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *sectionRepository) CountByCourseID(ctx context.Context, courseID int) (int, error) {
	return r.client.Sections.Query().
		Where(sections.CourseIDEQ(courseID)).
		Count(ctx)
}

func (r *sectionRepository) Update(ctx context.Context, id int, body dto.SectionUpdateReq) (*ent.Sections, error) {
	q := r.client.Sections.UpdateOneID(id)

	if body.Title != nil {
		q = q.SetTitle(*body.Title)
	}
	if body.SortOrder != nil {
		q = q.SetSortOrder(*body.SortOrder)
	}

	return q.Save(ctx)
}

func (r *sectionRepository) Delete(ctx context.Context, id int) error {
	return r.client.Sections.DeleteOneID(id).Exec(ctx)
}
