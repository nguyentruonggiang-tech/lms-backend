package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminUser(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/users")
		if err != nil {
			return err
		}
		op.SetTags("Admin User")
		op.SetSummary("Danh sách user")
		op.SetDescription("Có phân trang và filter theo role, status, keyword")
		op.AddReqStructure(new(struct {
			Page    int    `query:"page" example:"1"`
			Limit   int    `query:"limit" example:"10"`
			Role    string `query:"role" example:"student"`
			Status  string `query:"status" example:"active"`
			Keyword string `query:"keyword" example:"nguyen"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Users]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/users/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin User")
		op.SetSummary("Chi tiết user")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Users]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPatch, "/api/admin/users/{id}/status")
		if err != nil {
			return err
		}
		op.SetTags("Admin User")
		op.SetSummary("Khóa/mở user")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.UserUpdateStatusReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Users]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPatch, "/api/admin/users/{id}/role")
		if err != nil {
			return err
		}
		op.SetTags("Admin User")
		op.SetSummary("Đổi role")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.UserUpdateRoleReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Users]))
		reflector.AddOperation(op)
	}

	return nil
}
