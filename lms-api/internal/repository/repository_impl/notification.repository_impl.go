package repository_impl

import (
	"context"

	"lms-api/ent"
	"lms-api/ent/notifications"
	"lms-api/internal/common/pagination"
	"lms-api/internal/repository"
)

type notificationRepository struct {
	client *ent.Client
}

func NewNotificationRepository(client *ent.Client) repository.NotificationRepository {
	return &notificationRepository{client: client}
}

func (r *notificationRepository) FindAllByUserID(ctx context.Context, userID int, isRead *bool, query pagination.Query) ([]*ent.Notifications, error) {
	q := r.client.Notifications.Query().
		Where(notifications.UserIDEQ(userID)).
		Order(ent.Desc(notifications.FieldCreatedAt))

	if isRead != nil {
		q = q.Where(notifications.IsReadEQ(*isRead))
	}

	return q.Limit(query.Limit).Offset(query.Offset).All(ctx)
}

func (r *notificationRepository) CountByUserID(ctx context.Context, userID int, isRead *bool) (int, error) {
	q := r.client.Notifications.Query().
		Where(notifications.UserIDEQ(userID))

	if isRead != nil {
		q = q.Where(notifications.IsReadEQ(*isRead))
	}

	return q.Count(ctx)
}

func (r *notificationRepository) MarkRead(ctx context.Context, id, userID int) error {
	return r.client.Notifications.Update().
		Where(
			notifications.IDEQ(id),
			notifications.UserIDEQ(userID),
		).
		SetIsRead(true).
		Exec(ctx)
}

func (r *notificationRepository) MarkAllRead(ctx context.Context, userID int) error {
	return r.client.Notifications.Update().
		Where(notifications.UserIDEQ(userID)).
		SetIsRead(true).
		Exec(ctx)
}
