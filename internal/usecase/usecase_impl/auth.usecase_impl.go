package usecase_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
	"lms-backend/internal/usecase"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo     repository.UserRepository
	tokenUsecase usecase.TokenUsecase
}

func NewAuthUsecase(userRepo repository.UserRepository, tokenUsecase usecase.TokenUsecase) usecase.AuthUsecase {
	return &authUsecase{userRepo: userRepo, tokenUsecase: tokenUsecase}
}

func (a *authUsecase) Register(ctx context.Context, req dto.AuthRegisterReq) (*ent.Users, error) {
	exists, err := a.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	if exists {
		return nil, response.NewBadRequestException("Email already exists")
	}
	return a.userRepo.CreateUser(ctx, req)
}

func (a *authUsecase) Login(ctx context.Context, req dto.AuthLoginReq) (*dto.AuthLoginReturn, error) {
	user, err := a.userRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, response.NewBadRequestException("Invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, response.NewBadRequestException("Invalid email or password")
	}
	if user.Status == "blocked" {
		return nil, response.NewForbiddenException("Account is blocked")
	}
	return a.createTokenPair(user.ID, string(user.Role))
}

func (a *authUsecase) RefreshToken(ctx context.Context, accessToken, refreshToken string) (*dto.AuthRefreshTokenReturn, error) {
	claimAccess, err := a.tokenUsecase.VerifyAccessToken(accessToken, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}
	claimRefresh, err := a.tokenUsecase.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}
	if claimAccess.UserId != claimRefresh.UserId {
		return nil, response.NewUnauthorizedException("tokens do not belong to the same user")
	}
	user, err := a.userRepo.FindUserById(ctx, claimAccess.UserId)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}
	result, err := a.createTokenPair(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}
	return &dto.AuthRefreshTokenReturn{AccessToken: result.AccessToken, RefreshToken: result.RefreshToken}, nil
}

func (a *authUsecase) GetInfo(_ context.Context, user *ent.Users) (*ent.Users, error) {
	return user, nil
}

func (a *authUsecase) ChangePassword(ctx context.Context, userID int, req dto.AuthChangePasswordReq) error {
	user, err := a.userRepo.FindUserById(ctx, userID)
	if err != nil {
		return response.NewNotFoundException()
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return response.NewBadRequestException("Old password is incorrect")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return a.userRepo.UpdatePassword(ctx, userID, string(hashed))
}

func (a *authUsecase) createTokenPair(userID int, role string) (*dto.AuthLoginReturn, error) {
	access, err := a.tokenUsecase.CreateAccessToken(userID, role)
	if err != nil {
		return nil, err
	}
	refresh, err := a.tokenUsecase.CreateRefreshToken(userID, role)
	if err != nil {
		return nil, err
	}
	return &dto.AuthLoginReturn{AccessToken: access, RefreshToken: refresh}, nil
}
