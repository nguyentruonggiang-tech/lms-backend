package repository_impl

import (
	"context"

	"lms-worker/ent"
	"lms-worker/internal/repository"
)

type notificationRepository struct {
	client *ent.Client
}

func NewNotificationRepository(client *ent.Client) repository.NotificationRepository {
	return &notificationRepository{client: client}
}

func (r *notificationRepository) Create(ctx context.Context, userID int, title, content string) error {
	_, err := r.client.Notifications.Create().
		SetUserID(userID).
		SetTitle(title).
		SetContent(content).
		Save(ctx)
	return err
}
