package helpers

import (
	"errors"
	"lms-api/ent"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) (*ent.Users, error) {
	val, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("middleware protect not attached")
	}
	if val == nil {
		return nil, errors.New("user not found in context")
	}
	user, ok := val.(*ent.Users)
	if !ok {
		return nil, errors.New("invalid user type in context")
	}
	return user, nil
}
