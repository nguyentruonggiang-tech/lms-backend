package admin

import (
	"lms-api/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type DashboardDelivery struct {
	dashboardHandler *admin.DashboardHandler
}

func NewDashboardDelivery(dashboardHandler *admin.DashboardHandler) *DashboardDelivery {
	return &DashboardDelivery{dashboardHandler: dashboardHandler}
}

func (d *DashboardDelivery) RegisterRouter(adminGroup *gin.RouterGroup) {
	g := adminGroup.Group("dashboard")
	{
		g.GET("overview", d.dashboardHandler.GetOverview)
		g.GET("top-courses", d.dashboardHandler.GetTopCourses)
	}
}
