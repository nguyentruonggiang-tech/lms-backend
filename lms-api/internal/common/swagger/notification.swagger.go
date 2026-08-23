package swagger

import (
	"net/http"

	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"

	"github.com/swaggest/openapi-go/openapi3"
)

func notification(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/notifications")
		if err != nil {
			return err
		}
		op.SetTags("Notification")
		op.SetSummary("Danh sách thông báo của tôi (student only)")
		op.AddReqStructure(new(struct {
			IsRead *bool `query:"isRead" example:"false"`
			Page   int   `query:"page" example:"1"`
			Limit  int   `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Notifications]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPatch, "/api/my/notifications/{id}/read")
		if err != nil {
			return err
		}
		op.SetTags("Notification")
		op.SetSummary("Đánh dấu thông báo đã đọc")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPatch, "/api/my/notifications/read-all")
		if err != nil {
			return err
		}
		op.SetTags("Notification")
		op.SetSummary("Đánh dấu tất cả thông báo đã đọc")
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
