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

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) usecase.CategoryUsecase {
	return &categoryUsecase{categoryRepo: categoryRepo}
}

func (u *categoryUsecase) Create(ctx context.Context, body dto.CategoryCreateReq) (any, error) {
	data, err := u.categoryRepo.Create(ctx, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *categoryUsecase) FindAll(ctx context.Context, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.categoryRepo.FindAll(ctx, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.categoryRepo.Count(ctx)
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

func (u *categoryUsecase) FindByID(ctx context.Context, id int) (any, error) {
	data, err := u.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *categoryUsecase) Update(ctx context.Context, id int, body dto.CategoryUpdateReq) (any, error) {
	data, err := u.categoryRepo.Update(ctx, id, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *categoryUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := u.categoryRepo.Delete(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return true, nil
}
