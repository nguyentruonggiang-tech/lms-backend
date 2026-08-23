package usecase_impl

import (
	"context"
	"math"

	"lms-api/internal/common/cache"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type lessonUsecase struct {
	lessonRepo  repository.LessonRepository
	sectionRepo repository.SectionRepository
	redisClient *cache.RedisClient
}

func NewLessonUsecase(lessonRepo repository.LessonRepository, sectionRepo repository.SectionRepository, redisClient *cache.RedisClient) usecase.LessonUsecase {
	return &lessonUsecase{lessonRepo: lessonRepo, sectionRepo: sectionRepo, redisClient: redisClient}
}

func (u *lessonUsecase) Create(ctx context.Context, sectionID int, body dto.LessonCreateReq) (any, error) {
	section, err := u.sectionRepo.FindByID(ctx, sectionID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	data, err := u.lessonRepo.Create(ctx, sectionID, section.CourseID, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	u.redisClient.DelByPattern(ctx, "courses:*")
	return data, nil
}

func (u *lessonUsecase) FindByCourseID(ctx context.Context, courseID int, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.lessonRepo.FindByCourseID(ctx, courseID, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.lessonRepo.CountByCourseID(ctx, courseID)
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

func (u *lessonUsecase) FindByID(ctx context.Context, id int) (any, error) {
	data, err := u.lessonRepo.FindByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *lessonUsecase) Update(ctx context.Context, id int, body dto.LessonUpdateReq) (any, error) {
	data, err := u.lessonRepo.Update(ctx, id, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	u.redisClient.DelByPattern(ctx, "courses:*")
	return data, nil
}

func (u *lessonUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := u.lessonRepo.Delete(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	u.redisClient.DelByPattern(ctx, "courses:*")
	return true, nil
}
