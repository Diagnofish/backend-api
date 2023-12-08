package model

import (
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("secret-key")

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
