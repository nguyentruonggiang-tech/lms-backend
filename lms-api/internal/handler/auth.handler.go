package handler

import (
	"errors"
	"io"
	"lms-api/internal/common/helpers"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/usecase"
	"net/http"

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
	response.Success(result, "registered successfully", http.StatusCreated, ctx)
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
	setTokenCookie(ctx, result.AccessToken, result.RefreshToken)
	response.Success(nil, "login successful", 0, ctx)
}

func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("accessToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			return
		}
		ctx.Error(response.NewBadRequestException())
		return
	}
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			return
		}
		ctx.Error(response.NewBadRequestException())
		return
	}
	result, err := h.authUsecase.RefreshToken(ctx.Request.Context(), accessToken, refreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}
	setTokenCookie(ctx, result.AccessToken, result.RefreshToken)
	response.Success(nil, "token refreshed", 0, ctx)
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
	response.Success(result, "success", 0, ctx)
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
	response.Success(nil, "password changed", 0, ctx)
}

func setTokenCookie(ctx *gin.Context, accessToken, refreshToken string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("accessToken", accessToken, 0, "/", "", false, true)
	ctx.SetCookie("refreshToken", refreshToken, 0, "/", "", false, true)
}
