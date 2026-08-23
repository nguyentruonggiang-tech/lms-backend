package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"lms-worker/internal/common/elastic"
	"lms-worker/internal/usecase"
)

type SearchHandler struct {
	searchUsecase usecase.SearchUsecase
}

func NewSearchHandler(searchUsecase usecase.SearchUsecase) *SearchHandler {
	return &SearchHandler{searchUsecase: searchUsecase}
}

func (h *SearchHandler) HandleCourseIndex(ctx context.Context, body []byte) error {
	var doc elastic.CourseDoc
	if err := json.Unmarshal(body, &doc); err != nil {
		fmt.Printf("❌ [SEARCH] invalid payload: %v\n", err)
		return err
	}
	return h.searchUsecase.IndexCourse(ctx, doc)
}
