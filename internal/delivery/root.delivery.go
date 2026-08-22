package delivery

import "github.com/gin-gonic/gin"

type RootDelivery struct {
	authDelivery *authDelivery
}

func NewRootDelivery(authDelivery *authDelivery) *RootDelivery {
	return &RootDelivery{authDelivery: authDelivery}
}

func (r *RootDelivery) RegisterRouter(engine *gin.Engine) {
	apiGroup := engine.Group("api")
	{
		r.authDelivery.RegisterRouter(apiGroup)
	}
}
