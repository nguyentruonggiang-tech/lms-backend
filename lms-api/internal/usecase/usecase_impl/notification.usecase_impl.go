package usecase_impl

import (
	"context"
	"math"

	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type notificationUsecase struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationUsecase(notificationRepository repository.NotificationRepository) usecase.NotificationUsecase {
	return &notificationUsecase{notificationRepository: notificationRepository}
}

func (u *notificationUsecase) GetMyNotifications(ctx context.Context, userID int, isRead *bool, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.notificationRepository.FindAllByUserID(ctx, userID, isRead, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.notificationRepository.CountByUserID(ctx, userID, isRead)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return pagination.Response[any]{
		Items:     data,
		Page:      query.Page,
		Limit:     query.Limit,
		TotalItem: total,
		TotalPage: int(math.Ceil(float64(total) / float64(query.Limit))),
	}, nil
}

func (u *notificationUsecase) MarkRead(ctx context.Context, id, userID int) (any, error) {
	if err := u.notificationRepository.MarkRead(ctx, id, userID); err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return true, nil
}

func (u *notificationUsecase) MarkAllRead(ctx context.Context, userID int) (any, error) {
	if err := u.notificationRepository.MarkAllRead(ctx, userID); err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return true, nil
}
