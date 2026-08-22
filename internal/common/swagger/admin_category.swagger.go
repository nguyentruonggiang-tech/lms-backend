package swagger

import (
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
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
		op.SetSummary("Tạo category")
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
		op.SetSummary("Danh sách category")
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
		op.SetSummary("Chi tiết category")
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
		op.SetSummary("Cập nhật category")
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
		op.SetSummary("Xoá category")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
