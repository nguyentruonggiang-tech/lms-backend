package swagger

import (
	"net/http"

	"lms-api/ent"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"

	"github.com/swaggest/openapi-go/openapi3"
)

func lessonProgress(reflector *openapi3.Reflector) error {
	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/courses/{courseId}/lessons")
		if err != nil {
			return err
		}
		op.SetTags("Lesson Progress")
		op.SetSummary("Danh sách bài học trong khóa đã đăng ký")
		op.AddReqStructure(new(struct {
			CourseId string `path:"courseId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[[]dto.LessonWithProgress]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/lessons/{lessonId}")
		if err != nil {
			return err
		}
		op.SetTags("Lesson Progress")
		op.SetSummary("Xem chi tiết bài học")
		op.AddReqStructure(new(struct {
			LessonId string `path:"lessonId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[dto.LessonWithProgress]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodPost, "/api/my/lessons/{lessonId}/complete")
		if err != nil {
			return err
		}
		op.SetTags("Lesson Progress")
		op.SetSummary("Đánh dấu hoàn thành bài học")
		op.AddReqStructure(new(struct {
			LessonId string `path:"lessonId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[*ent.LessonProgresses]))
		reflector.AddOperation(op)
	}

	{
		op, err := reflector.NewOperationContext(http.MethodGet, "/api/my/courses/{courseId}/progress")
		if err != nil {
			return err
		}
		op.SetTags("Lesson Progress")
		op.SetSummary("Xem tiến độ khóa học")
		op.AddReqStructure(new(struct {
			CourseId string `path:"courseId" example:"1"`
		}))
		op.AddRespStructure(new(response.SuccessFormat[[]*ent.LessonProgresses]))
		reflector.AddOperation(op)
	}

	return nil
}
