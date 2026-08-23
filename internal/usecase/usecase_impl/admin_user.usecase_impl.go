package usecase_impl

import (
	"context"
	"math"

	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
	"lms-backend/internal/usecase"
)

type adminUserUsecase struct {
	userRepo repository.UserRepository
}

func NewAdminUserUsecase(userRepo repository.UserRepository) usecase.AdminUserUsecase {
	return &adminUserUsecase{userRepo: userRepo}
}

func (u *adminUserUsecase) FindAll(ctx context.Context, page, limit, role, status, keyword string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.userRepo.FindAll(ctx, query, role, status, keyword)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.userRepo.CountAll(ctx, role, status, keyword)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return pagination.Response[any]{
		Items:     data,
		Page:      query.Page,
		Limit:     query.Limit,
		TotalItem: total,
		TotalPage: int(math.Ceil(float64(total) / float64(query.Limit))),
	}, nil
}

func (u *adminUserUsecase) FindByID(ctx context.Context, id int) (any, error) {
	data, err := u.userRepo.FindUserById(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *adminUserUsecase) UpdateStatus(ctx context.Context, id int, body dto.UserUpdateStatusReq) (any, error) {
	data, err := u.userRepo.UpdateStatus(ctx, id, body.Status)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *adminUserUsecase) UpdateRole(ctx context.Context, id int, body dto.UserUpdateRoleReq) (any, error) {
	data, err := u.userRepo.UpdateRole(ctx, id, body.Role)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}
