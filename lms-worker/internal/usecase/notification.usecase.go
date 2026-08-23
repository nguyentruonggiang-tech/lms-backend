package usecase

import "context"

type NotificationUsecase interface {
	CreateCourseEnrolled(ctx context.Context, userID int, title, content string) error
}
