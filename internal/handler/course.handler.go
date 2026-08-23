package handler

import (
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseUsecase usecase.CourseUsecase
}

func NewCourseHandler(courseUsecase usecase.CourseUsecase) *CourseHandler {
	return &CourseHandler{courseUsecase: courseUsecase}
}

func (h *CourseHandler) FindAllPublished(ctx *gin.Context) {
	filter := dto.CoursePublicFilter{
		Level: ctx.Query("level"),
	}

	if v := ctx.Query("categoryId"); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			filter.CategoryID = &id
		}
	}
	if v := ctx.Query("minPrice"); v != "" {
		if p, err := strconv.ParseFloat(v, 64); err == nil {
			filter.MinPrice = &p
		}
	}
	if v := ctx.Query("maxPrice"); v != "" {
		if p, err := strconv.ParseFloat(v, 64); err == nil {
			filter.MaxPrice = &p
		}
	}

	data, err := h.courseUsecase.FindAllPublished(ctx.Request.Context(), filter, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *CourseHandler) Search(ctx *gin.Context) {
	filter := dto.CoursePublicFilter{
		Q:     ctx.Query("q"),
		Level: ctx.Query("level"),
	}

	if v := ctx.Query("categoryId"); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			filter.CategoryID = &id
		}
	}

	data, err := h.courseUsecase.SearchPublished(ctx.Request.Context(), filter, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *CourseHandler) FindPublishedByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err2 := h.courseUsecase.FindPublishedByID(ctx.Request.Context(), id)
	if err2 != nil {
		ctx.Error(err2)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *CourseHandler) FindPreviewLessons(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err2 := h.courseUsecase.FindPreviewLessons(ctx.Request.Context(), id)
	if err2 != nil {
		ctx.Error(err2)
		return
	}
	response.Success(data, "success", 0, ctx)
}
