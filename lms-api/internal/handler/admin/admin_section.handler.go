package admin

import (
	"strconv"

	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SectionHandler struct {
	sectionUsecase usecase.SectionUsecase
}

func NewSectionHandler(sectionUsecase usecase.SectionUsecase) *SectionHandler {
	return &SectionHandler{sectionUsecase: sectionUsecase}
}

func (h *SectionHandler) Create(ctx *gin.Context) {
	courseID, err := strconv.Atoi(ctx.Param("courseId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid courseId"))
		return
	}

	var body dto.SectionCreateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.sectionUsecase.Create(ctx.Request.Context(), courseID, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *SectionHandler) FindByCourseID(ctx *gin.Context) {
	courseID, err := strconv.Atoi(ctx.Param("courseId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid courseId"))
		return
	}

	data, err := h.sectionUsecase.FindByCourseID(ctx.Request.Context(), courseID, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *SectionHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("courseId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.SectionUpdateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.sectionUsecase.Update(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *SectionHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("courseId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.sectionUsecase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}
