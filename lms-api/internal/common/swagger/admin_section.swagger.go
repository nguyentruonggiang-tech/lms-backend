package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminSection(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/courses/{courseId}/sections")
		if err != nil {
			return err
		}
		op.SetTags("Admin Section")
		op.SetSummary("Tạo chương")
		op.AddReqStructure(new(struct {
			CourseID string `path:"courseId" example:"1"`
		}))
		op.AddReqStructure(new(dto.SectionCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Sections]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/courses/{courseId}/sections")
		if err != nil {
			return err
		}
		op.SetTags("Admin Section")
		op.SetSummary("Danh sách chương")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			CourseID string `path:"courseId" example:"1"`
			Page     int    `query:"page" example:"1"`
			Limit    int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Sections]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPut, "/api/admin/sections/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Section")
		op.SetSummary("Cập nhật chương")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.SectionUpdateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Sections]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/admin/sections/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Section")
		op.SetSummary("Xóa mềm chương")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
