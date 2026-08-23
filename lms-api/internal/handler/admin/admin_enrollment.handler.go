package admin

import (
	"strconv"

	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	adminEnrollmentUsecase usecase.AdminEnrollmentUsecase
}

func NewEnrollmentHandler(adminEnrollmentUsecase usecase.AdminEnrollmentUsecase) *EnrollmentHandler {
	return &EnrollmentHandler{adminEnrollmentUsecase: adminEnrollmentUsecase}
}

func (h *EnrollmentHandler) FindAll(ctx *gin.Context) {
	filter := dto.AdminEnrollmentFilter{
		Status: ctx.Query("status"),
	}

	if v := ctx.Query("courseId"); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			filter.CourseID = &id
		}
	}
	if v := ctx.Query("userId"); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			filter.UserID = &id
		}
	}

	data, err := h.adminEnrollmentUsecase.FindAll(ctx.Request.Context(), filter, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *EnrollmentHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.adminEnrollmentUsecase.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *EnrollmentHandler) UpdateStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.EnrollmentUpdateStatusReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.adminEnrollmentUsecase.UpdateStatus(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}
