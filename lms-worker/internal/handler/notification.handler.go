package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"lms-worker/internal/usecase"
)

type NotificationHandler struct {
	notificationUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{notificationUsecase: notificationUsecase}
}

type courseEnrolledPayload struct {
	UserID  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *NotificationHandler) HandleCourseEnrolled(ctx context.Context, body []byte) error {
	var payload courseEnrolledPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Printf("❌ [NOTIFICATION] invalid payload: %v\n", err)
		return err
	}
	return h.notificationUsecase.CreateCourseEnrolled(ctx, payload.UserID, payload.Title, payload.Content)
}

func (h *NotificationHandler) HandleCertificateIssued(ctx context.Context, body []byte) error {
	var payload courseEnrolledPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Printf("❌ [NOTIFICATION] invalid payload: %v\n", err)
		return err
	}
	return h.notificationUsecase.CreateCourseEnrolled(ctx, payload.UserID, payload.Title, payload.Content)
}
