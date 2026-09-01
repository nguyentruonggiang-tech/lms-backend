package swagger

import (
	"lms-api/ent"
	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func adminLesson(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/admin/sections/{sectionId}/lessons")
		if err != nil {
			return err
		}
		op.SetTags("Admin Lesson")
		op.SetSummary("Tạo bài học")
		op.AddReqStructure(new(struct {
			SectionID string `path:"sectionId" example:"1"`
		}))
		op.AddReqStructure(new(dto.LessonCreateReq))
		op.AddRespStructure(new(response.SuccessFormat[*ent.Lessons]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/admin/courses/{courseId}/lessons")
		if err != nil {
			return err
		}
		op.SetTags("Admin Lesson")
		op.SetSummary("Danh sách bài học của khóa")
		op.SetDescription("Có phân trang")
		op.AddReqStructure(new(struct {
			CourseID string `path:"courseId" example:"1"`
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
		op.SetSummary("Xóa mềm bài học")
		op.AddReqStructure(new(struct {
			Id string `path:"id" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[bool]))
		reflector.AddOperation(op)
	}

	return nil
}
