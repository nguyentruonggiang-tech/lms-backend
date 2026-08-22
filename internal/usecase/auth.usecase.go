package usecase

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/dto"
)

type AuthUsecase interface {
	Register(ctx context.Context, req dto.AuthRegisterReq) (*ent.Users, error)
	Login(ctx context.Context, req dto.AuthLoginReq) (*dto.AuthLoginReturn, error)
	RefreshToken(ctx context.Context, accessToken, refreshToken string) (*dto.AuthRefreshTokenReturn, error)
	GetInfo(ctx context.Context, user *ent.Users) (*ent.Users, error)
	ChangePassword(ctx context.Context, userID int, req dto.AuthChangePasswordReq) error
}
