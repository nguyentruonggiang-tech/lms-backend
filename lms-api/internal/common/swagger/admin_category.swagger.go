package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminCategory(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/categories")
		if err != nil {
			return err
		}
		op.SetTags("Admin Category")
		op.SetSummary("Tạo danh mục")
		op.AddReqStructure(new(dto.CategoryCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Categories]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/categories")
		if err != nil {
			return err
		}
		op.SetTags("Admin Category")
		op.SetSummary("Danh sách danh mục")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			Page  int `query:"page" example:"1"`
			Limit int `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Categories]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/categories/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Category")
		op.SetSummary("Chi tiết danh mục")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Categories]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPut, "/api/admin/categories/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Category")
		op.SetSummary("Cập nhật danh mục")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.CategoryUpdateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Categories]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/admin/categories/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Category")
		op.SetSummary("Xóa mềm danh mục")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
