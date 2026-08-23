package swagger

import (
	"net/http"

	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"

	"github.com/swaggest/openapi-go/openapi3"
)

func quizClient(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/quizzes/{quizId}")
		if err != nil {
			return err
		}
		op.SetTags("Quiz Client")
		op.SetSummary("Lấy quiz để làm bài (ẩn correct_option)")
		op.AddReqStructure(new(struct {
			QuizId string `path:"quizId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[dto.QuizWithQuestionsRes]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/my/quizzes/{quizId}/submit")
		if err != nil {
			return err
		}
		op.SetTags("Quiz Client")
		op.SetSummary("Nộp bài quiz")
		op.AddReqStructure(new(struct {
			QuizId  string             `path:"quizId" example:"1"`
			Answers []dto.QuizAnswerItem `json:"answers"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.QuizAttempts]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/quizzes/{quizId}/attempts")
		if err != nil {
			return err
		}
		op.SetTags("Quiz Client")
		op.SetSummary("Lịch sử làm quiz")
		op.AddReqStructure(new(struct {
			QuizId string `path:"quizId" example:"1"`
			Page   int    `query:"page" example:"1"`
			Limit  int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.QuizAttempts]]))
		reflector.AddOperation(op)
	}

	return nil
}
