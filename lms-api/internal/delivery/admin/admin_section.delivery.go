package admin

import (
	"lms-api/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type SectionDelivery struct {
	sectionHandler *admin.SectionHandler
}

func NewSectionDelivery(sectionHandler *admin.SectionHandler) *SectionDelivery {
	return &SectionDelivery{sectionHandler: sectionHandler}
}

func (d *SectionDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	courses := adminGroup.Group("courses/:courseId/sections")
	{
		courses.POST("", d.sectionHandler.Create)
		courses.GET("", d.sectionHandler.FindByCourseID)
	}

	sections := adminGroup.Group("sections")
	{
		sections.PUT(":id", d.sectionHandler.Update)
		sections.DELETE(":id", d.sectionHandler.Delete)
	}
}
