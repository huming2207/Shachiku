package common

import "github.com/dgrijalva/jwt-go"

type JwtUserClaims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}
