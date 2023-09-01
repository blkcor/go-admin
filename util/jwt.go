package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var SecretKey = "blkcor"

func GenerateJWT(issuer string) (string, error) {
	//claims = payload + signature
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	//generate the token
	return claims.SignedString([]byte(SecretKey))
}

func ParseJWT(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims
	issuer, _ := claims.GetIssuer()
	return issuer, nil
}
