package swagger

import (
	"net/http"

	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"

	"github.com/swaggest/openapi-go/openapi3"
)

func enrollment(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/enrollments")
		if err != nil {
			return err
		}
		op.SetTags("Enrollment")
		op.SetSummary("Đăng ký học (student only)")
		op.AddReqStructure(new(struct {
			CourseID int `json:"courseId" required:"true" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Enrollments]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/enrollments")
		if err != nil {
			return err
		}
		op.SetTags("Enrollment")
		op.SetSummary("Danh sách khóa học đã đăng ký")
		op.AddReqStructure(new(struct {
			Status string `query:"status" example:"active"`
			Page   int    `query:"page" example:"1"`
			Limit  int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Enrollments]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/courses/{courseId}")
		if err != nil {
			return err
		}
		op.SetTags("Enrollment")
		op.SetSummary("Chi tiết khóa học đã đăng ký")
		op.AddReqStructure(new(struct {
			CourseId string `path:"courseId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Enrollments]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/my/enrollments/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Enrollment")
		op.SetSummary("Hủy đăng ký học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
