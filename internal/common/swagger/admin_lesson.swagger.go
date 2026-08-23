package swagger

import (
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminLesson(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/sections/{section_id}/lessons")
		if err != nil {
			return err
		}
		op.SetTags("Admin Lesson")
		op.SetSummary("Tạo bài học")
		op.AddReqStructure(new(struct {
			SectionID string `path:"section_id" example:"1"`
		}))
		op.AddReqStructure(new(dto.LessonCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Lessons]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/courses/{course_id}/lessons")
		if err != nil {
			return err
		}
		op.SetTags("Admin Lesson")
		op.SetSummary("Danh sách bài học theo course")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			CourseID string `path:"course_id" example:"1"`
			Page     int    `query:"page" example:"1"`
			Limit    int    `query:"limit" example:"10"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[pagination.Response[*ent.Lessons]]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/lessons/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Lesson")
		op.SetSummary("Chi tiết bài học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Lessons]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPut, "/api/admin/lessons/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Lesson")
		op.SetSummary("Cập nhật bài học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddReqStructure(new(dto.LessonUpdateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Lessons]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodDelete, "/api/admin/lessons/{id}")
		if err != nil {
			return err
		}
		op.SetTags("Admin Lesson")
		op.SetSummary("Xoá bài học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
