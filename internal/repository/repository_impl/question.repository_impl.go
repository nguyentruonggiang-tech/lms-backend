package repository_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/ent/questions"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
)

type questionRepository struct {
	client *ent.Client
}

func NewQuestionRepository(client *ent.Client) repository.QuestionRepository {
	return &questionRepository{client: client}
}

func (r *questionRepository) Create(ctx context.Context, quizID int, body dto.QuestionCreateReq) (*ent.Questions, error) {
	return r.client.Questions.Create().
		SetQuizID(quizID).
		SetQuestionText(body.QuestionText).
		SetOptionA(body.OptionA).
		SetOptionB(body.OptionB).
		SetOptionC(body.OptionC).
		SetOptionD(body.OptionD).
		SetCorrectOption(questions.CorrectOption(body.CorrectOption)).
		Save(ctx)
}

func (r *questionRepository) FindByQuizID(ctx context.Context, quizID int, query pagination.Query) ([]*ent.Questions, error) {
	return r.client.Questions.Query().
		Where(questions.QuizIDEQ(quizID)).
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *questionRepository) CountByQuizID(ctx context.Context, quizID int) (int, error) {
	return r.client.Questions.Query().
		Where(questions.QuizIDEQ(quizID)).
		Count(ctx)
}

func (r *questionRepository) Update(ctx context.Context, id int, body dto.QuestionUpdateReq) (*ent.Questions, error) {
	q := r.client.Questions.UpdateOneID(id)

	if body.QuestionText != nil {
		q = q.SetQuestionText(*body.QuestionText)
	}
	if body.OptionA != nil {
		q = q.SetOptionA(*body.OptionA)
	}
	if body.OptionB != nil {
		q = q.SetOptionB(*body.OptionB)
	}
	if body.OptionC != nil {
		q = q.SetOptionC(*body.OptionC)
	}
	if body.OptionD != nil {
		q = q.SetOptionD(*body.OptionD)
	}
	if body.CorrectOption != nil {
		q = q.SetCorrectOption(questions.CorrectOption(*body.CorrectOption))
	}

	return q.Save(ctx)
}

func (r *questionRepository) Delete(ctx context.Context, id int) error {
	return r.client.Questions.DeleteOneID(id).Exec(ctx)
}
