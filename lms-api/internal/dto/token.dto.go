package dto

import "github.com/golang-jwt/jwt/v5"

type CustomClaim struct {
	UserId int
	Role   string
	jwt.RegisteredClaims
}
