package usecase

import "context"

type LessonProgressUsecase interface {
	ListLessons(ctx context.Context, userID, courseID int) (any, error)
	FindLesson(ctx context.Context, userID, lessonID int) (any, error)
	Complete(ctx context.Context, userID, lessonID int) (any, error)
	GetProgress(ctx context.Context, userID, courseID int) (any, error)
}
