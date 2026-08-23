package repository

import "context"

type NotificationRepository interface {
	Create(ctx context.Context, userID int, title, content string) error
}
