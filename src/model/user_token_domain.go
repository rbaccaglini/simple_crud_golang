package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {

	secret := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("error trying to generate token. err=%s", err.Error()))
	}
	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv("JWT_SECRET_KEY")

	t, err := jwt.Parse(removeBearerPrefix(tokenValue), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil
}

func removeBearerPrefix(token string) string {
	var PREFIX = "Bearer "
	return strings.TrimPrefix(token, PREFIX)
}
