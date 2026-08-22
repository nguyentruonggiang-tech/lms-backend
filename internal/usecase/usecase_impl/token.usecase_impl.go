package usecase_impl

import (
	"errors"
	"lms-backend/internal/common/env"
	"lms-backend/internal/dto"
	"lms-backend/internal/usecase"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenUsecase struct {
	env *env.Env
}

func NewTokenUsecase(e *env.Env) usecase.TokenUsecase {
	return &tokenUsecase{env: e}
}

func (t *tokenUsecase) CreateAccessToken(userId int, role string) (string, error) {
	claim := dto.CustomClaim{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.env.ExpiresAtAccessToken)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(t.env.SecretAccessToken))
}

func (t *tokenUsecase) CreateRefreshToken(userId int, role string) (string, error) {
	claim := dto.CustomClaim{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.env.ExpiresAtRefreshToken)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(t.env.SecretRefreshToken))
}

func (t *tokenUsecase) VerifyAccessToken(token string, options ...jwt.ParserOption) (*dto.CustomClaim, error) {
	claim := &dto.CustomClaim{}
	parsed, err := jwt.ParseWithClaims(token, claim, func(*jwt.Token) (any, error) {
		return []byte(t.env.SecretAccessToken), nil
	}, options...)
	if err != nil {
		return nil, err
	}
	if !parsed.Valid {
		return nil, errors.New("access token invalid")
	}
	return claim, nil
}

func (t *tokenUsecase) VerifyRefreshToken(token string, options ...jwt.ParserOption) (*dto.CustomClaim, error) {
	claim := &dto.CustomClaim{}
	parsed, err := jwt.ParseWithClaims(token, claim, func(*jwt.Token) (any, error) {
		return []byte(t.env.SecretRefreshToken), nil
	}, options...)
	if err != nil {
		return nil, err
	}
	if !parsed.Valid {
		return nil, errors.New("refresh token invalid")
	}
	return claim, nil
}
