package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	user_response "github.com/rbaccaglini/simple_crud_golang/internal/models/response/user"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
)

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetName() string
	GetAge() int8
	GetPassword() string

	SetID(string)

	EncryptPassword()
	TokenGenerate() (string, *rest_err.RestErr)
	ConvertDomainToResponse() user_response.UserResponse
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(uid, name string, age int8) UserDomainInterface {
	return &userDomain{
		id:   uid,
		name: name,
		age:  age,
	}
}

func NewLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) ConvertDomainToResponse() user_response.UserResponse {
	return user_response.UserResponse{
		ID:    ud.GetID(),
		Email: ud.GetEmail(),
		Name:  ud.GetName(),
		Age:   ud.GetAge(),
	}
}

func (ud *userDomain) TokenGenerate() (string, *rest_err.RestErr) {
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
