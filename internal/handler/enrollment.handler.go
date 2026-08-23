package handler

import (
	"strconv"

	"lms-backend/internal/common/helpers"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	enrollmentUsecase usecase.EnrollmentUsecase
}

func NewEnrollmentHandler(enrollmentUsecase usecase.EnrollmentUsecase) *EnrollmentHandler {
	return &EnrollmentHandler{enrollmentUsecase: enrollmentUsecase}
}

func (h *EnrollmentHandler) Enroll(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	var body dto.EnrollReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.enrollmentUsecase.Enroll(ctx.Request.Context(), user.ID, body)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "enrolled successfully", 201, ctx)
}

func (h *EnrollmentHandler) FindMyEnrollments(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	data, err := h.enrollmentUsecase.FindMyEnrollments(ctx.Request.Context(), user.ID, ctx.Query("status"), ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *EnrollmentHandler) FindMyCourse(ctx *gin.Context) {
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

	data, err := h.enrollmentUsecase.FindMyCourse(ctx.Request.Context(), user.ID, courseID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *EnrollmentHandler) Cancel(ctx *gin.Context) {
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

	data, err := h.enrollmentUsecase.Cancel(ctx.Request.Context(), user.ID, id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "cancelled successfully", 0, ctx)
}
