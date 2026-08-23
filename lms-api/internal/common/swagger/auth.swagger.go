package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func auth(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/auth/register")
		if err != nil {
			return err
		}
		op.SetTags("Auth")
		op.SetSummary("Đăng ký tài khoản")
		op.AddReqStructure(new(dto.AuthRegisterReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Users]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/auth/login")
		if err != nil {
			return err
		}
		op.SetTags("Auth")
		op.SetSummary("Đăng nhập")
		op.AddReqStructure(new(dto.AuthLoginReq))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/auth/refresh-token")
		if err != nil {
			return err
		}
		op.SetTags("Auth")
		op.SetSummary("Làm mới access token")
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/auth/get-info")
		if err != nil {
			return err
		}
		op.SetTags("Auth")
		op.SetSummary("Lấy thông tin user hiện tại")
		op.AddRespStructure(new(response.SuccessFormat[*ent.Users]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPatch, "/api/auth/change-password")
		if err != nil {
			return err
		}
		op.SetTags("Auth")
		op.SetSummary("Đổi mật khẩu")
		op.AddReqStructure(new(dto.AuthChangePasswordReq))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
