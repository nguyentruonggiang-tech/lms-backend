package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"lms-api/internal/common/helpers"
	"lms-api/internal/common/response"
	"lms-api/internal/usecase"
)

type NotificationHandler struct {
	notificationUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{notificationUsecase: notificationUsecase}
}

func (h *NotificationHandler) GetMyNotifications(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	var isRead *bool
	if val := ctx.Query("isRead"); val != "" {
		b, err := strconv.ParseBool(val)
		if err == nil {
			isRead = &b
		}
	}

	data, err := h.notificationUsecase.GetMyNotifications(ctx.Request.Context(), user.ID, isRead, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "Success", 0, ctx)
}

func (h *NotificationHandler) MarkRead(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.notificationUsecase.MarkRead(ctx.Request.Context(), id, user.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "Success", 0, ctx)
}

func (h *NotificationHandler) MarkAllRead(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	data, err := h.notificationUsecase.MarkAllRead(ctx.Request.Context(), user.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "Success", 0, ctx)
}
