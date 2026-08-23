package repository_impl

import (
	"context"
	"time"

	"lms-backend/ent"
	"lms-backend/ent/lessonprogresses"
	"lms-backend/internal/repository"
)

type lessonProgressRepository struct {
	client *ent.Client
}

func NewLessonProgressRepository(client *ent.Client) repository.LessonProgressRepository {
	return &lessonProgressRepository{client: client}
}

func (r *lessonProgressRepository) FindByUserAndLesson(ctx context.Context, userID, lessonID int) (*ent.LessonProgresses, error) {
	return r.client.LessonProgresses.Query().
		Where(
			lessonprogresses.UserIDEQ(userID),
			lessonprogresses.LessonIDEQ(lessonID),
		).
		Only(ctx)
}

func (r *lessonProgressRepository) Create(ctx context.Context, userID, courseID, lessonID int) (*ent.LessonProgresses, error) {
	now := time.Now()
	return r.client.LessonProgresses.Create().
		SetUserID(userID).
		SetCourseID(courseID).
		SetLessonID(lessonID).
		SetIsCompleted(true).
		SetCompletedAt(now).
		Save(ctx)
}

func (r *lessonProgressRepository) MarkComplete(ctx context.Context, id int) (*ent.LessonProgresses, error) {
	return r.client.LessonProgresses.UpdateOneID(id).
		SetIsCompleted(true).
		SetCompletedAt(time.Now()).
		Save(ctx)
}

func (r *lessonProgressRepository) CountCompletedByUserAndCourse(ctx context.Context, userID, courseID int) (int, error) {
	return r.client.LessonProgresses.Query().
		Where(
			lessonprogresses.UserIDEQ(userID),
			lessonprogresses.CourseIDEQ(courseID),
			lessonprogresses.IsCompletedEQ(true),
		).
		Count(ctx)
}

func (r *lessonProgressRepository) FindAllByUserAndCourse(ctx context.Context, userID, courseID int) ([]*ent.LessonProgresses, error) {
	return r.client.LessonProgresses.Query().
		WithLessons().
		Where(
			lessonprogresses.UserIDEQ(userID),
			lessonprogresses.CourseIDEQ(courseID),
		).
		All(ctx)
}
