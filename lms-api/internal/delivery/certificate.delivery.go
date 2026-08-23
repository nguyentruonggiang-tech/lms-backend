package delivery

import (
	"lms-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type certificateDelivery struct {
	certificateHandler *handler.CertificateHandler
}

func NewCertificateDelivery(certificateHandler *handler.CertificateHandler) *certificateDelivery {
	return &certificateDelivery{certificateHandler: certificateHandler}
}

func (d *certificateDelivery) RegisterPublicRouter(apiGroup *gin.RouterGroup) {
	apiGroup.GET("certificates/:code", d.certificateHandler.GetByCode)
}

func (d *certificateDelivery) RegisterStudentRouter(studentGroup *gin.RouterGroup) {
	myGroup := studentGroup.Group("my")
	myGroup.GET("certificates", d.certificateHandler.GetMyCertificates)
}
