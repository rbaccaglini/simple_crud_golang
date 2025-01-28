package service

import (
	"errors"
	"os"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_LoginUserService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := mocks.NewMockUserRepository(ctl)
	srv := &userDomainService{r}

	t.Run("user not found", func(t *testing.T) {

		userId := primitive.NewObjectID().String()
		email := "test@test.com"
		password := "123$%¨7"
		ud := model.NewUserDomain(email, password, "Test Silva", 20)
		ud.SetID(userId)

		r.EXPECT().
			FindUserByEmailAndPass(gomock.Any(), gomock.Any()).
			Return(nil, rest_err.NewNotFoundError("User not found."))

		rud, tk, err := srv.LoginUserService(ud)

		assert.Nil(t, rud)
		assert.Empty(t, tk)
		assert.EqualValues(t, 403, err.Code)

	})

	t.Run("user not found", func(t *testing.T) {

		userId := primitive.NewObjectID().String()
		email := "test@test.com"
		password := "123$%¨7"
		ud := model.NewUserDomain(email, password, "Test Silva", 20)
		ud.SetID(userId)

		r.EXPECT().
			FindUserByEmailAndPass(gomock.Any(), gomock.Any()).
			Return(nil, rest_err.NewInternalServerError("error"))

		rud, tk, err := srv.LoginUserService(ud)

		assert.Nil(t, rud)
		assert.Empty(t, tk)
		assert.EqualValues(t, 500, err.Code)

	})

	t.Run("error with token", func(t *testing.T) {

		udMock := mocks.NewMockUserDomainInterface(ctl)
		udMock.EXPECT().GenerateToken().Return("", rest_err.NewInternalServerError("error"))
		udMock.EXPECT().GetEmail().Return("test@test.com")
		udMock.EXPECT().GetPassword().Return("123$%¨7")
		udMock.EXPECT().EncryptPassword()

		r.EXPECT().
			FindUserByEmailAndPass(gomock.Any(), gomock.Any()).
			Return(udMock, nil)

		rud, tk, err := srv.LoginUserService(udMock)

		assert.Nil(t, rud)
		assert.Empty(t, tk)
		assert.EqualValues(t, 500, err.Code)

	})

	t.Run("success", func(t *testing.T) {

		os.Setenv("JWT_SECRET_KEY", "123456")
		defer os.Clearenv()

		userId := primitive.NewObjectID().String()
		email := "test@test.com"
		password := "123$%¨7"
		ud := model.NewUserDomain(email, password, "Test Silva", 20)
		ud.SetID(userId)

		r.EXPECT().
			FindUserByEmailAndPass(gomock.Any(), gomock.Any()).
			Return(ud, nil)

		rud, tk, err := srv.LoginUserService(ud)

		assert.Nil(t, err)
		assert.NotEmpty(t, tk)
		assert.EqualValues(t, ud, rud)
		assert.Nil(t, isTokenValid(tk))
	})
}

func isTokenValid(token string) error {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		}
		return nil, errors.New("invalid token")
	})
	if err != nil {
		return errors.New("invalid token")
	}

	_, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return errors.New("invalid token")
	}
	return nil
}
