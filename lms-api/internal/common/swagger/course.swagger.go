package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func course(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/courses")
		if err != nil {
			return err
		}
		op.SetTags("Course")
		op.SetSummary("Danh sách khóa học đã publish")
		op.AddReqStructure(new(struct {
			Page       int     `query:"page" example:"1"`
			Limit      int     `query:"limit" example:"10"`
			CategoryID int     `query:"categoryId" example:"1"`
			Level      string  `query:"level" example:"beginner"`
			MinPrice   float64 `query:"minPrice" example:"0"`
			MaxPrice   float64 `query:"maxPrice" example:"500000"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Courses]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/courses/search")
		if err != nil {
			return err
		}
		op.SetTags("Course")
		op.SetSummary("Tìm kiếm khóa học")
		op.SetDescription("Tìm theo title và description (MySQL LIKE)")
		op.AddReqStructure(new(struct {
			Q          string `query:"q" example:"golang"`
			CategoryID int    `query:"categoryId" example:"1"`
			Level      string `query:"level" example:"beginner"`
			Page       int    `query:"page" example:"1"`
			Limit      int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Courses]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/courses/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Course")
		op.SetSummary("Chi tiết khóa học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Courses]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/courses/{id}/preview-lessons")
		if err != nil {
			return err
		}
		op.SetTags("Course")
		op.SetSummary("Danh sách bài học xem thử")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[[]*ent.Lessons]))
		reflector.AddOperation(op)
	}

	return nil
}
