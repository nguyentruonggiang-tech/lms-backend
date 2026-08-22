package delivery

import "github.com/gin-gonic/gin"

type RootDelivery struct{}

func NewRootDelivery() *RootDelivery {
	return &RootDelivery{}
}

func (r *RootDelivery) RegisterRouter(engine *gin.Engine) {
	engine.Group("api")
}
