package handler

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"lms-api/internal/common/helpers"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var body dto.AuthRegisterReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		if errors.Is(err, io.EOF) {
			ctx.Error(response.NewBadRequestException("body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}
	result, err := h.authUsecase.Register(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(result, "Registered successfully", http.StatusCreated, ctx)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var body dto.AuthLoginReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		if errors.Is(err, io.EOF) {
			ctx.Error(response.NewBadRequestException("body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}
	result, err := h.authUsecase.Login(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(result, "Login successful", 0, ctx)
}

func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	var body dto.AuthRefreshTokenReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	accessToken := ""
	authHeader := ctx.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		accessToken = strings.TrimPrefix(authHeader, "Bearer ")
	}

	result, err := h.authUsecase.RefreshToken(ctx.Request.Context(), accessToken, body.RefreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(result, "Token refreshed", 0, ctx)
}

func (h *AuthHandler) GetInfo(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}
	result, err := h.authUsecase.GetInfo(ctx.Request.Context(), user)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(result, "Success", 0, ctx)
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}
	var body dto.AuthChangePasswordReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}
	if err := h.authUsecase.ChangePassword(ctx.Request.Context(), user.ID, body); err != nil {
		ctx.Error(err)
		return
	}
	response.Success(nil, "Password changed", 0, ctx)
}
