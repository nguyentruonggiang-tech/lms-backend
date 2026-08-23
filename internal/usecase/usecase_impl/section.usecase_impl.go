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

type sectionUsecase struct {
	sectionRepo repository.SectionRepository
}

func NewSectionUsecase(sectionRepo repository.SectionRepository) usecase.SectionUsecase {
	return &sectionUsecase{sectionRepo: sectionRepo}
}

func (u *sectionUsecase) Create(ctx context.Context, courseID int, body dto.SectionCreateReq) (any, error) {
	data, err := u.sectionRepo.Create(ctx, courseID, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *sectionUsecase) FindByCourseID(ctx context.Context, courseID int, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.sectionRepo.FindByCourseID(ctx, courseID, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.sectionRepo.CountByCourseID(ctx, courseID)
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

func (u *sectionUsecase) Update(ctx context.Context, id int, body dto.SectionUpdateReq) (any, error) {
	data, err := u.sectionRepo.Update(ctx, id, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *sectionUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := u.sectionRepo.Delete(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return true, nil
}
