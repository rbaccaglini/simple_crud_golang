package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
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

func VerifyToken(c *gin.Context) {
	secret := os.Getenv("JWT_SECRET_KEY")

	tokenValue := c.GetHeader("Authorization")

	t, err := jwt.Parse(removeBearerPrefix(tokenValue), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})
	if err != nil {
		re := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(re.Code, re)
		c.Abort()
		return
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		re := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(re.Code, re)
		c.Abort()
		return
	}

	u := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", u))
}

func removeBearerPrefix(token string) string {
	var PREFIX = "Bearer "
	return strings.TrimPrefix(token, PREFIX)
}
