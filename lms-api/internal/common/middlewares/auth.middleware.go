package middlewares

import (
	"errors"
	"lms-api/ent/users"
	"lms-api/internal/common/helpers"
	"lms-api/internal/common/response"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	tokenUsecase usecase.TokenUsecase
	userRepo     repository.UserRepository
}

func NewAuthMiddleware(tokenUsecase usecase.TokenUsecase, userRepo repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{tokenUsecase: tokenUsecase, userRepo: userRepo}
}

func (a *AuthMiddleware) Protect(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("accessToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			ctx.Abort()
			return
		}
		ctx.Error(response.NewBadRequestException())
		ctx.Abort()
		return
	}

	claim, err := a.tokenUsecase.VerifyAccessToken(accessToken)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			ctx.Error(response.NewForbiddenException("token expired"))
			ctx.Abort()
			return
		}
		ctx.Error(response.NewUnauthorizedException())
		ctx.Abort()
		return
	}

	user, err := a.userRepo.FindUserById(ctx, claim.UserId)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}

func (a *AuthMiddleware) AdminOnly(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		ctx.Abort()
		return
	}

	if user.Role != users.RoleAdmin {
		ctx.Error(response.NewForbiddenException())
		ctx.Abort()
		return
	}

	ctx.Next()
}
