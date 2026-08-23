package usecase

import (
	"context"

	"lms-worker/internal/common/elastic"
)

type SearchUsecase interface {
	IndexCourse(ctx context.Context, doc elastic.CourseDoc) error
}
