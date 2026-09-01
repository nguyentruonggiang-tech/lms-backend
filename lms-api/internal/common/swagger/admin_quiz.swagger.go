package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminQuiz(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/lessons/{lessonId}/quizzes")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Tạo quiz")
		op.AddReqStructure(new(struct {
			LessonID string `path:"lessonId" example:"1"`
		}))
		op.AddReqStructure(new(dto.QuizCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Quizzes]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/courses/{courseId}/quizzes")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Danh sách quiz của khóa")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			CourseID string `path:"courseId" example:"1"`
			Page     int    `query:"page" example:"1"`
			Limit    int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Quizzes]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/quizzes/{quizId}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Chi tiết quiz")
		op.AddReqStructure(new(struct {
			Id string `path:"quizId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Quizzes]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPut, "/api/admin/quizzes/{quizId}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Cập nhật quiz")
		op.AddReqStructure(new(struct {
			Id string `path:"quizId" example:"1"`
		}))
		op.AddReqStructure(new(dto.QuizUpdateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Quizzes]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/admin/quizzes/{quizId}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Quiz")
		op.SetSummary("Xóa mềm quiz")
		op.AddReqStructure(new(struct {
			Id string `path:"quizId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
