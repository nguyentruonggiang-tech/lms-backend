package di

import (
	"lms-backend/ent"
	"lms-backend/internal/common/env"

	"github.com/gin-gonic/gin"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, e *env.Env) {
}
