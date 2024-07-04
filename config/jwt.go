package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("akjsdafnihgfwur8943ufuiw2f")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
