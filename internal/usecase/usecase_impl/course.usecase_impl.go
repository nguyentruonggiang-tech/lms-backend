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
	lessonRepository repository.LessonRepository
}

func NewCourseUsecase(courseRepository repository.CourseRepository, lessonRepository repository.LessonRepository) usecase.CourseUsecase {
	return &courseUsecase{courseRepository: courseRepository, lessonRepository: lessonRepository}
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

func (u *courseUsecase) UpdateStatus(ctx context.Context, id int, body dto.CourseUpdateStatusReq) (any, error) {
	data, err := u.courseRepository.UpdateStatus(ctx, id, body.Status)
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

func (u *courseUsecase) FindAllPublished(ctx context.Context, filter dto.CoursePublicFilter, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.courseRepository.FindAllPublished(ctx, filter, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.courseRepository.CountPublished(ctx, filter)
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

func (u *courseUsecase) SearchPublished(ctx context.Context, filter dto.CoursePublicFilter, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.courseRepository.SearchPublished(ctx, filter, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.courseRepository.CountSearch(ctx, filter)
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

func (u *courseUsecase) FindPublishedByID(ctx context.Context, id int) (any, error) {
	data, err := u.courseRepository.FindPublishedByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *courseUsecase) FindPreviewLessons(ctx context.Context, courseID int) (any, error) {
	_, err := u.courseRepository.FindPublishedByID(ctx, courseID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	data, err := u.lessonRepository.FindPreviewByCourseID(ctx, courseID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}
