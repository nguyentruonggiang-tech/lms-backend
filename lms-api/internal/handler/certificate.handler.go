package handler

import (
	"github.com/gin-gonic/gin"

	"lms-api/internal/common/helpers"
	"lms-api/internal/common/response"
	"lms-api/internal/usecase"
)

type CertificateHandler struct {
	certificateUsecase usecase.CertificateUsecase
}

func NewCertificateHandler(certificateUsecase usecase.CertificateUsecase) *CertificateHandler {
	return &CertificateHandler{certificateUsecase: certificateUsecase}
}

func (h *CertificateHandler) GetMyCertificates(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	data, err := h.certificateUsecase.GetMyCertificates(ctx.Request.Context(), user.ID, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *CertificateHandler) GetByCode(ctx *gin.Context) {
	data, err := h.certificateUsecase.GetByCode(ctx.Request.Context(), ctx.Param("code"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}
