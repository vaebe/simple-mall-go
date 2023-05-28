package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	ID          int32
	NickName    string
	AuthorityId int32
	jwt.RegisteredClaims
}
