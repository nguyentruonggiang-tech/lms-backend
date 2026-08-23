package usecase_impl

import (
	"context"
	"math"

	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type questionUsecase struct {
	questionRepo repository.QuestionRepository
}

func NewQuestionUsecase(questionRepo repository.QuestionRepository) usecase.QuestionUsecase {
	return &questionUsecase{questionRepo: questionRepo}
}

func (u *questionUsecase) Create(ctx context.Context, quizID int, body dto.QuestionCreateReq) (any, error) {
	data, err := u.questionRepo.Create(ctx, quizID, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *questionUsecase) FindByQuizID(ctx context.Context, quizID int, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.questionRepo.FindByQuizID(ctx, quizID, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.questionRepo.CountByQuizID(ctx, quizID)
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

func (u *questionUsecase) Update(ctx context.Context, id int, body dto.QuestionUpdateReq) (any, error) {
	data, err := u.questionRepo.Update(ctx, id, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *questionUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := u.questionRepo.Delete(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return true, nil
}
