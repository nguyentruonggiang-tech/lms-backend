package swagger

import (
	"net/http"

	"lms-api/internal/common/response"
	"lms-api/internal/dto"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminDashboard(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/dashboard/overview")
		if err != nil {
			return err
		}
		op.SetTags("Admin Dashboard")
		op.SetSummary("Tổng quan số học viên, khóa học, enrollment, chứng chỉ")
		op.AddRespStructure(new(response.SuccessFormat[dto.DashboardOverview]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/dashboard/top-courses")
		if err != nil {
			return err
		}
		op.SetTags("Admin Dashboard")
		op.SetSummary("Top khóa học nhiều người học nhất")
		op.AddReqStructure(new(struct {
			FromDate string `query:"fromDate" example:"2026-01-01"`
			ToDate   string `query:"toDate" example:"2026-12-31"`
			Limit    int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[[]dto.TopCourseItem]))
		reflector.AddOperation(op)
	}

	return nil
}
