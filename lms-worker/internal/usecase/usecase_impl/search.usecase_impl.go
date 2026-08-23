package usecase_impl

import (
	"context"

	"lms-worker/internal/common/elastic"
	"lms-worker/internal/usecase"
)

type searchUsecase struct {
	esClient *elastic.ElasticClient
}

func NewSearchUsecase(esClient *elastic.ElasticClient) usecase.SearchUsecase {
	return &searchUsecase{esClient: esClient}
}

func (u *searchUsecase) IndexCourse(ctx context.Context, doc elastic.CourseDoc) error {
	return u.esClient.IndexCourse(ctx, doc)
}
