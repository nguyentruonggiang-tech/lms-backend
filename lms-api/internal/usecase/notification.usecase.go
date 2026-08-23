package usecase

import "context"

type NotificationUsecase interface {
	GetMyNotifications(ctx context.Context, userID int, isRead *bool, page, limit string) (any, error)
	MarkRead(ctx context.Context, id, userID int) (any, error)
	MarkAllRead(ctx context.Context, userID int) (any, error)
}
