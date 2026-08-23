package repository

import (
	"context"

	"lms-api/ent"
	"lms-api/internal/common/pagination"
)

type NotificationRepository interface {
	FindAllByUserID(ctx context.Context, userID int, isRead *bool, query pagination.Query) ([]*ent.Notifications, error)
	CountByUserID(ctx context.Context, userID int, isRead *bool) (int, error)
	MarkRead(ctx context.Context, id, userID int) error
	MarkAllRead(ctx context.Context, userID int) error
}
