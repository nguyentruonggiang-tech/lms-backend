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
		op.SetSummary("Tạo khóa học")
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
		op.SetSummary("Danh sách khóa học")
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
		op.SetSummary("Chi tiết khóa học")
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
		op.SetSummary("Cập nhật khóa học")
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
		op.SetSummary("Đổi trạng thái khóa học")
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
		op.SetSummary("Xóa mềm khóa học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/courses/{id}/reindex")
		if err != nil {
			return err
		}
		op.SetTags("Admin Course")
		op.SetSummary("Đẩy lại khóa học vào Elasticsearch")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
