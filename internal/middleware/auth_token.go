package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
)

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

	u := domain.NewUserDomain(
		claims["email"].(string),
		"",
		claims["name"].(string),
		int8(claims["age"].(float64)),
	)
	u.SetID(claims["id"].(string))

	// u := &userDomain{
	// 	id:    claims["id"].(string),
	// 	email: claims["email"].(string),
	// 	name:  claims["name"].(string),
	// 	age:   int8(claims["age"].(float64)),
	// }

	logger.Info(fmt.Sprintf("User authenticated: %#v", u))
}

func removeBearerPrefix(token string) string {
	var PREFIX = "Bearer "
	return strings.TrimPrefix(token, PREFIX)
}
