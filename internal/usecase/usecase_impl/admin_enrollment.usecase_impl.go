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

type adminEnrollmentUsecase struct {
	enrollmentRepository repository.EnrollmentRepository
}

func NewAdminEnrollmentUsecase(enrollmentRepository repository.EnrollmentRepository) usecase.AdminEnrollmentUsecase {
	return &adminEnrollmentUsecase{enrollmentRepository: enrollmentRepository}
}

func (u *adminEnrollmentUsecase) FindAll(ctx context.Context, filter dto.AdminEnrollmentFilter, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.enrollmentRepository.FindAll(ctx, filter, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.enrollmentRepository.CountAll(ctx, filter)
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

func (u *adminEnrollmentUsecase) FindByID(ctx context.Context, id int) (any, error) {
	data, err := u.enrollmentRepository.FindByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *adminEnrollmentUsecase) UpdateStatus(ctx context.Context, id int, body dto.EnrollmentUpdateStatusReq) (any, error) {
	_, err := u.enrollmentRepository.FindByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	data, err := u.enrollmentRepository.UpdateStatus(ctx, id, body.Status)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}
