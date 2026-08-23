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

type quizUsecase struct {
	quizRepo   repository.QuizRepository
	lessonRepo repository.LessonRepository
}

func NewQuizUsecase(quizRepo repository.QuizRepository, lessonRepo repository.LessonRepository) usecase.QuizUsecase {
	return &quizUsecase{quizRepo: quizRepo, lessonRepo: lessonRepo}
}

func (u *quizUsecase) Create(ctx context.Context, lessonID int, body dto.QuizCreateReq) (any, error) {
	lesson, err := u.lessonRepo.FindByID(ctx, lessonID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	data, err := u.quizRepo.Create(ctx, lessonID, lesson.CourseID, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *quizUsecase) FindByLessonID(ctx context.Context, lessonID int, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.quizRepo.FindByLessonID(ctx, lessonID, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.quizRepo.CountByLessonID(ctx, lessonID)
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

func (u *quizUsecase) FindByID(ctx context.Context, id int) (any, error) {
	data, err := u.quizRepo.FindByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *quizUsecase) Update(ctx context.Context, id int, body dto.QuizUpdateReq) (any, error) {
	data, err := u.quizRepo.Update(ctx, id, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *quizUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := u.quizRepo.Delete(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return true, nil
}
