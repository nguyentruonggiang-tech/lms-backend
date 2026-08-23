package usecase_impl

import (
	"context"

	"lms-worker/internal/repository"
	"lms-worker/internal/usecase"
)

type notificationUsecase struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationUsecase(notificationRepository repository.NotificationRepository) usecase.NotificationUsecase {
	return &notificationUsecase{notificationRepository: notificationRepository}
}

func (u *notificationUsecase) CreateCourseEnrolled(ctx context.Context, userID int, title, content string) error {
	return u.notificationRepository.Create(ctx, userID, title, content)
}
