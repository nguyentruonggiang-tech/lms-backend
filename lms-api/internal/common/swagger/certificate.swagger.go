package swagger

import (
	"net/http"

	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"

	"github.com/swaggest/openapi-go/openapi3"
)

func certificate(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/certificates")
		if err != nil {
			return err
		}
		op.SetTags("Certificate")
		op.SetSummary("Danh sách chứng chỉ của tôi (student only)")
		op.AddReqStructure(new(struct {
			Page  int `query:"page" example:"1"`
			Limit int `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Certificates]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/certificates/{code}")
		if err != nil {
			return err
		}
		op.SetTags("Certificate")
		op.SetSummary("Xem chứng chỉ public bằng code")
		op.AddReqStructure(new(struct {
			Code string `path:"code" example:"550e8400-e29b-41d4-a716-446655440000"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Certificates]))
		reflector.AddOperation(op)
	}

	return nil
}
