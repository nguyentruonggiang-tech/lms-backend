package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminCourse(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/courses")
		if err != nil {
			return err
		}
		op.SetTags("Admin Course")
		op.SetSummary("Tạo course")
		op.AddReqStructure(new(dto.CourseCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Courses]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/courses")
		if err != nil {
			return err
		}
		op.SetTags("Admin Course")
		op.SetSummary("Danh sách course")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			Page  int `query:"page" example:"1"`
			Limit int `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Courses]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/courses/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Course")
		op.SetSummary("Chi tiết course")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Courses]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPut, "/api/admin/courses/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Course")
		op.SetSummary("Cập nhật course")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.CourseUpdateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Courses]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPatch, "/api/admin/courses/{id}/status")
		if err != nil {
			return err
		}
		op.SetTags("Admin Course")
		op.SetSummary("Đổi trạng thái course")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.CourseUpdateStatusReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Courses]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/admin/courses/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Course")
		op.SetSummary("Xoá course")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
