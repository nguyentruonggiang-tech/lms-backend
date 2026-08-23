package swagger

import (
	"net/http"

	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminEnrollment(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/enrollments")
		if err != nil {
			return err
		}
		op.SetTags("Admin Enrollment")
		op.SetSummary("Danh sách đăng ký học")
		op.AddReqStructure(new(struct {
			CourseID int    `query:"courseId" example:"1"`
			UserID   int    `query:"userId" example:"1"`
			Status   string `query:"status" example:"active"`
			Page     int    `query:"page" example:"1"`
			Limit    int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Enrollments]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/enrollments/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Enrollment")
		op.SetSummary("Chi tiết đăng ký học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Enrollments]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPatch, "/api/admin/enrollments/{id}/status")
		if err != nil {
			return err
		}
		op.SetTags("Admin Enrollment")
		op.SetSummary("Đổi trạng thái đăng ký học")
		op.AddReqStructure(new(struct {
			Id     string `path:"id" example:"1"`
			Status string `json:"status" example:"active"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Enrollments]))
		reflector.AddOperation(op)
	}

	return nil
}
