package repository

import (
	"context"
	"lms-backend/ent"
)

type LessonProgressRepository interface {
	FindByUserAndLesson(ctx context.Context, userID, lessonID int) (*ent.LessonProgresses, error)
	Create(ctx context.Context, userID, courseID, lessonID int) (*ent.LessonProgresses, error)
	MarkComplete(ctx context.Context, id int) (*ent.LessonProgresses, error)
	CountCompletedByUserAndCourse(ctx context.Context, userID, courseID int) (int, error)
	FindAllByUserAndCourse(ctx context.Context, userID, courseID int) ([]*ent.LessonProgresses, error)
}
