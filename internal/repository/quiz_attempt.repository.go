package repository

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
)

type QuizAttemptRepository interface {
	Create(ctx context.Context, userID, quizID, total, correct int, score float64, isPassed bool) (*ent.QuizAttempts, error)
	FindByUserAndQuiz(ctx context.Context, userID, quizID int, query pagination.Query) ([]*ent.QuizAttempts, error)
	CountByUserAndQuiz(ctx context.Context, userID, quizID int) (int, error)
}
