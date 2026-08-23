package admin

import (
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardUsecase usecase.DashboardUsecase
}

func NewDashboardHandler(dashboardUsecase usecase.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{dashboardUsecase: dashboardUsecase}
}

func (h *DashboardHandler) GetOverview(ctx *gin.Context) {
	data, err := h.dashboardUsecase.GetOverview(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *DashboardHandler) GetTopCourses(ctx *gin.Context) {
	var req dto.TopCourseReq
	_ = ctx.ShouldBindQuery(&req)

	data, err := h.dashboardUsecase.GetTopCourses(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}
