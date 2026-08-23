package usecase

import (
	"lms-api/internal/dto"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUsecase interface {
	CreateAccessToken(userId int, role string) (string, error)
	CreateRefreshToken(userId int, role string) (string, error)
	VerifyAccessToken(token string, options ...jwt.ParserOption) (*dto.CustomClaim, error)
	VerifyRefreshToken(token string, options ...jwt.ParserOption) (*dto.CustomClaim, error)
}
