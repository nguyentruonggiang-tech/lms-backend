package repository_impl

import (
	"context"
	"lms-api/ent"
	"lms-api/ent/lessons"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
	"lms-api/internal/repository"
)

type lessonRepository struct {
	client *ent.Client
}

func NewLessonRepository(client *ent.Client) repository.LessonRepository {
	return &lessonRepository{client: client}
}

func (r *lessonRepository) Create(ctx context.Context, sectionID, courseID int, body dto.LessonCreateReq) (*ent.Lessons, error) {
	return r.client.Lessons.Create().
		SetSectionID(sectionID).
		SetCourseID(courseID).
		SetTitle(body.Title).
		SetNillableContent(body.Content).
		SetNillableVideoURL(body.VideoURL).
		SetNillableDurationMinutes(body.DurationMinutes).
		SetNillableSortOrder(body.SortOrder).
		SetNillableIsPreview(body.IsPreview).
		Save(ctx)
}

func (r *lessonRepository) FindByCourseID(ctx context.Context, courseID int, query pagination.Query) ([]*ent.Lessons, error) {
	return r.client.Lessons.Query().
		Where(lessons.CourseIDEQ(courseID)).
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *lessonRepository) CountByCourseID(ctx context.Context, courseID int) (int, error) {
	return r.client.Lessons.Query().
		Where(lessons.CourseIDEQ(courseID)).
		Count(ctx)
}

func (r *lessonRepository) FindByID(ctx context.Context, id int) (*ent.Lessons, error) {
	return r.client.Lessons.Get(ctx, id)
}

func (r *lessonRepository) Update(ctx context.Context, id int, body dto.LessonUpdateReq) (*ent.Lessons, error) {
	q := r.client.Lessons.UpdateOneID(id)

	if body.Title != nil {
		q = q.SetTitle(*body.Title)
	}
	if body.Content != nil {
		q = q.SetContent(*body.Content)
	}
	if body.VideoURL != nil {
		q = q.SetVideoURL(*body.VideoURL)
	}
	if body.DurationMinutes != nil {
		q = q.SetDurationMinutes(*body.DurationMinutes)
	}
	if body.SortOrder != nil {
		q = q.SetSortOrder(*body.SortOrder)
	}
	if body.IsPreview != nil {
		q = q.SetIsPreview(*body.IsPreview)
	}

	return q.Save(ctx)
}

func (r *lessonRepository) Delete(ctx context.Context, id int) error {
	return r.client.Lessons.DeleteOneID(id).Exec(ctx)
}

func (r *lessonRepository) FindPreviewByCourseID(ctx context.Context, courseID int) ([]*ent.Lessons, error) {
	return r.client.Lessons.Query().
		Where(
			lessons.CourseIDEQ(courseID),
			lessons.IsPreviewEQ(true),
		).
		Order(lessons.BySortOrder()).
		All(ctx)
}
