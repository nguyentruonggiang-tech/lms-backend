package repository_impl

import (
	"context"
	"lms-api/ent"
	"lms-api/ent/quizzes"
	"lms-api/internal/common/pagination"
	"lms-api/internal/dto"
	"lms-api/internal/repository"
)

type quizRepository struct {
	client *ent.Client
}

func NewQuizRepository(client *ent.Client) repository.QuizRepository {
	return &quizRepository{client: client}
}

func (r *quizRepository) Create(ctx context.Context, lessonID, courseID int, body dto.QuizCreateReq) (*ent.Quizzes, error) {
	q := r.client.Quizzes.Create().
		SetLessonID(lessonID).
		SetCourseID(courseID).
		SetTitle(body.Title)

	if body.PassingScore != nil {
		q = q.SetPassingScore(*body.PassingScore)
	}

	return q.Save(ctx)
}

func (r *quizRepository) FindByLessonID(ctx context.Context, lessonID int, query pagination.Query) ([]*ent.Quizzes, error) {
	return r.client.Quizzes.Query().
		Where(quizzes.LessonIDEQ(lessonID)).
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *quizRepository) CountByLessonID(ctx context.Context, lessonID int) (int, error) {
	return r.client.Quizzes.Query().
		Where(quizzes.LessonIDEQ(lessonID)).
		Count(ctx)
}

func (r *quizRepository) FindByID(ctx context.Context, id int) (*ent.Quizzes, error) {
	return r.client.Quizzes.Get(ctx, id)
}

func (r *quizRepository) Update(ctx context.Context, id int, body dto.QuizUpdateReq) (*ent.Quizzes, error) {
	q := r.client.Quizzes.UpdateOneID(id)

	if body.Title != nil {
		q = q.SetTitle(*body.Title)
	}
	if body.PassingScore != nil {
		q = q.SetPassingScore(*body.PassingScore)
	}

	return q.Save(ctx)
}

func (r *quizRepository) Delete(ctx context.Context, id int) error {
	return r.client.Quizzes.DeleteOneID(id).Exec(ctx)
}
