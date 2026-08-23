package repository_impl

import (
	"context"

	"lms-api/ent"
	"lms-api/ent/quizattempts"
	"lms-api/internal/common/pagination"
	"lms-api/internal/repository"
)

type quizAttemptRepository struct {
	client *ent.Client
}

func NewQuizAttemptRepository(client *ent.Client) repository.QuizAttemptRepository {
	return &quizAttemptRepository{client: client}
}

func (r *quizAttemptRepository) Create(ctx context.Context, userID, quizID, total, correct int, score float64, isPassed bool) (*ent.QuizAttempts, error) {
	return r.client.QuizAttempts.Create().
		SetUserID(userID).
		SetQuizID(quizID).
		SetTotalQuestions(total).
		SetCorrectAnswers(correct).
		SetScore(score).
		SetIsPassed(isPassed).
		Save(ctx)
}

func (r *quizAttemptRepository) FindByUserAndQuiz(ctx context.Context, userID, quizID int, query pagination.Query) ([]*ent.QuizAttempts, error) {
	return r.client.QuizAttempts.Query().
		Where(
			quizattempts.UserIDEQ(userID),
			quizattempts.QuizIDEQ(quizID),
		).
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *quizAttemptRepository) CountByUserAndQuiz(ctx context.Context, userID, quizID int) (int, error) {
	return r.client.QuizAttempts.Query().
		Where(
			quizattempts.UserIDEQ(userID),
			quizattempts.QuizIDEQ(quizID),
		).
		Count(ctx)
}
