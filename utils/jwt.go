package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
