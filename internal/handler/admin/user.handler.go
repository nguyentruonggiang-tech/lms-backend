package admin

import (
	"strconv"

	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	adminUserUsecase usecase.AdminUserUsecase
}

func NewUserHandler(adminUserUsecase usecase.AdminUserUsecase) *UserHandler {
	return &UserHandler{adminUserUsecase: adminUserUsecase}
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	data, err := h.adminUserUsecase.FindAll(
		ctx.Request.Context(),
		ctx.Query("page"),
		ctx.Query("limit"),
		ctx.Query("role"),
		ctx.Query("status"),
		ctx.Query("keyword"),
	)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *UserHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.adminUserUsecase.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *UserHandler) UpdateStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.UserUpdateStatusReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.adminUserUsecase.UpdateStatus(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *UserHandler) UpdateRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.UserUpdateRoleReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.adminUserUsecase.UpdateRole(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}
