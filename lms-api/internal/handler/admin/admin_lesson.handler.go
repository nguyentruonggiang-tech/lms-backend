package admin

import (
	"strconv"

	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	lessonUsecase usecase.LessonUsecase
}

func NewLessonHandler(lessonUsecase usecase.LessonUsecase) *LessonHandler {
	return &LessonHandler{lessonUsecase: lessonUsecase}
}

func (h *LessonHandler) Create(ctx *gin.Context) {
	sectionID, err := strconv.Atoi(ctx.Param("section_id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid section_id"))
		return
	}

	var body dto.LessonCreateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.lessonUsecase.Create(ctx.Request.Context(), sectionID, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *LessonHandler) FindByCourseID(ctx *gin.Context) {
	courseID, err := strconv.Atoi(ctx.Param("course_id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid course_id"))
		return
	}

	data, err := h.lessonUsecase.FindByCourseID(ctx.Request.Context(), courseID, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *LessonHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.lessonUsecase.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *LessonHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.LessonUpdateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.lessonUsecase.Update(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *LessonHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.lessonUsecase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}
