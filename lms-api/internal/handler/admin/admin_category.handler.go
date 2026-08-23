package admin

import (
	"strconv"

	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase}
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var body dto.CategoryCreateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.categoryUsecase.Create(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *CategoryHandler) FindAll(ctx *gin.Context) {
	data, err := h.categoryUsecase.FindAll(ctx.Request.Context(), ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *CategoryHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.categoryUsecase.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *CategoryHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.CategoryUpdateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.categoryUsecase.Update(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *CategoryHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.categoryUsecase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}
