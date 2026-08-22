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

type courseUsecase struct {
	courseRepository repository.CourseRepository
}

func NewCourseUsecase(courseRepository repository.CourseRepository) usecase.CourseUsecase {
	return &courseUsecase{courseRepository: courseRepository}
}

func (u *courseUsecase) Create(ctx context.Context, body dto.CourseCreateReq) (any, error) {
	data, err := u.courseRepository.Create(ctx, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *courseUsecase) FindAll(ctx context.Context, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.courseRepository.FindAll(ctx, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.courseRepository.Count(ctx)
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

func (u *courseUsecase) FindByID(ctx context.Context, id int) (any, error) {
	data, err := u.courseRepository.FindByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *courseUsecase) Update(ctx context.Context, id int, body dto.CourseUpdateReq) (any, error) {
	data, err := u.courseRepository.Update(ctx, id, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *courseUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := u.courseRepository.Delete(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return true, nil
}
