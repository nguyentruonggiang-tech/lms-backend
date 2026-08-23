package handler

import (
	"strconv"

	"lms-backend/internal/common/helpers"
	"lms-backend/internal/common/response"
	"lms-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type LessonProgressHandler struct {
	lessonProgressUsecase usecase.LessonProgressUsecase
}

func NewLessonProgressHandler(lessonProgressUsecase usecase.LessonProgressUsecase) *LessonProgressHandler {
	return &LessonProgressHandler{lessonProgressUsecase: lessonProgressUsecase}
}

func (h *LessonProgressHandler) ListLessons(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	courseID, err := strconv.Atoi(ctx.Param("courseId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid courseId"))
		return
	}

	data, err := h.lessonProgressUsecase.ListLessons(ctx.Request.Context(), user.ID, courseID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *LessonProgressHandler) FindLesson(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	lessonID, err := strconv.Atoi(ctx.Param("lessonId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid lessonId"))
		return
	}

	data, err := h.lessonProgressUsecase.FindLesson(ctx.Request.Context(), user.ID, lessonID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *LessonProgressHandler) Complete(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	lessonID, err := strconv.Atoi(ctx.Param("lessonId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid lessonId"))
		return
	}

	data, err := h.lessonProgressUsecase.Complete(ctx.Request.Context(), user.ID, lessonID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *LessonProgressHandler) GetProgress(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	courseID, err := strconv.Atoi(ctx.Param("courseId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid courseId"))
		return
	}

	data, err := h.lessonProgressUsecase.GetProgress(ctx.Request.Context(), user.ID, courseID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}
