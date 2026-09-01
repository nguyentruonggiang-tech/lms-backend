package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminQuestion(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/quizzes/{quizId}/questions")
		if err != nil {
			return err
		}
		op.SetTags("Admin Question")
		op.SetSummary("Tạo câu hỏi")
		op.AddReqStructure(new(struct {
			QuizID string `path:"quizId" example:"1"`
		}))
		op.AddReqStructure(new(dto.QuestionCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Questions]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/quizzes/{quizId}/questions")
		if err != nil {
			return err
		}
		op.SetTags("Admin Question")
		op.SetSummary("Danh sách câu hỏi")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			QuizID string `path:"quizId" example:"1"`
			Page   int    `query:"page" example:"1"`
			Limit  int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Questions]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPut, "/api/admin/questions/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Question")
		op.SetSummary("Cập nhật câu hỏi")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.QuestionUpdateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Questions]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/admin/questions/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Question")
		op.SetSummary("Xóa mềm câu hỏi")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
