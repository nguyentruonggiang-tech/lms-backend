package swagger

import (
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminQuiz(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/lessons/{lesson_id}/quizzes")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Tạo quiz")
		op.AddReqStructure(new(struct {
			LessonID string `path:"lesson_id" example:"1"`
		}))
		op.AddReqStructure(new(dto.QuizCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Quizzes]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/lessons/{lesson_id}/quizzes")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Danh sách quiz theo lesson")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			LessonID string `path:"lesson_id" example:"1"`
			Page     int    `query:"page" example:"1"`
			Limit    int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Quizzes]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/quizzes/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Chi tiết quiz")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Quizzes]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPut, "/api/admin/quizzes/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Cập nhật quiz")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.QuizUpdateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Quizzes]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/admin/quizzes/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Xoá quiz")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
