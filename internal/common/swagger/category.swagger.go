package swagger

import (
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func category(reflector *openapi3.Reflector) error {
	op, err := reflector.NewOperationContext(http.MethodGet, "/api/categories")
	if err != nil {
		return err
	}
	op.SetTags("Category")
	op.SetSummary("Danh sách danh mục")
	op.AddReqStructure(new(struct {
		Page  int `query:"page" example:"1"`
		Limit int `query:"limit" example:"10"`
	}))
	op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Categories]]))
	reflector.AddOperation(op)
	return nil
}
