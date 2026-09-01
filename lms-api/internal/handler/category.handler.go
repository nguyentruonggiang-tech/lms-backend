package handler

import (
	"lms-api/internal/common/response"
	"lms-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase}
}

func (h *CategoryHandler) FindAll(ctx *gin.Context) {
	data, err := h.categoryUsecase.FindAll(ctx.Request.Context(), ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "Success", 0, ctx)
}
