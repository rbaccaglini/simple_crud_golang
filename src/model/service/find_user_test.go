package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_FindUserByIdService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := mocks.NewMockUserRepository(ctl)
	srv := NewUserDomainService(r)

	id := "12345"
	mud := model.NewUserDomain(
		"test@test.com",
		"123$%¨7",
		"Test Silva",
		22,
	)
	mud.SetID(id)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		r.EXPECT().FindUserById(id).Return(mud, nil)

		ud, err := srv.FindUserByIdService(id)

		assert.Nil(t, err)
		assert.EqualValues(t, mud, ud)
	})

	t.Run("when_user_not_found", func(t *testing.T) {
		r.EXPECT().FindUserById(id).Return(
			nil,
			rest_err.NewNotFoundError("user not found"),
		)

		ud, err := srv.FindUserByIdService(id)

		assert.Nil(t, ud)
		assert.EqualValues(t, err.Code, 404)
	})
}

func TestUserDomainService_FindUserByEmailService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	id := "12345"
	email := "test@test.com"

	r := mocks.NewMockUserRepository(ctl)
	srv := NewUserDomainService(r)
	mud := model.NewUserDomain(
		email,
		"123$%¨7",
		"Test Silva",
		22,
	)
	mud.SetID(id)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		r.EXPECT().FindUserByEmail(email).Return(mud, nil)

		ud, err := srv.FindUserByEmailService(email)

		assert.Nil(t, err)
		assert.EqualValues(t, mud, ud)
	})

	t.Run("when_user_not_found", func(t *testing.T) {
		r.EXPECT().FindUserByEmail(email).Return(
			nil,
			rest_err.NewNotFoundError("User not found."),
		)

		ud, err := srv.FindUserByEmailService(email)

		assert.Nil(t, ud)
		assert.EqualValues(t, err.Code, 404)
	})

	t.Run("when_user_not_found", func(t *testing.T) {
		r.EXPECT().FindUserByEmail(email).Return(
			nil,
			rest_err.NewInternalServerError("Error trying to find user"),
		)

		ud, err := srv.FindUserByEmailService(email)

		assert.Nil(t, ud)
		assert.EqualValues(t, err.Code, 500)
	})
}

func TestUserDomainService_FindAllUsersService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := mocks.NewMockUserRepository(ctl)
	srv := NewUserDomainService(r)

	t.Run("when_exists_an_user_list_returns_success", func(t *testing.T) {

		mud1 := model.NewUserDomain(
			"test1@test.com",
			"123$%¨7",
			"Test1 Silva",
			21,
		)
		mud1.SetID(primitive.NewObjectID().String())

		mud2 := model.NewUserDomain(
			"test2@test.com",
			"223$%¨7",
			"Test2 Silva",
			22,
		)
		mud2.SetID(primitive.NewObjectID().String())

		muds := []model.UserDomainInterface{mud1, mud2}

		r.EXPECT().FindAllUsers().Return(muds, nil)

		uds, err := srv.FindAllUsersService()

		assert.Nil(t, err)
		assert.EqualValues(t, len(muds), len(uds))
	})

	t.Run("when_user_not_found", func(t *testing.T) {
		r.EXPECT().FindAllUsers().Return(
			nil,
			rest_err.NewNotFoundError("User not found."),
		)

		ud, err := srv.FindAllUsersService()

		assert.Nil(t, ud)
		assert.EqualValues(t, err.Code, 404)
	})

	t.Run("when_internal_error", func(t *testing.T) {
		r.EXPECT().FindAllUsers().Return(
			nil,
			rest_err.NewInternalServerError("User not found."),
		)

		ud, err := srv.FindAllUsersService()

		assert.Nil(t, ud)
		assert.EqualValues(t, err.Code, 500)
	})
}

func TestUserDomainService_findUserByEmailAndPassService(t *testing.T) {
	id := "12345"
	email := "test@test.com"
	password := "123$%¨7"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mocks.NewMockUserRepository(ctrl)
	srv := &userDomainService{r}
	mud := model.NewUserDomain(
		email,
		password,
		"Test Silva",
		22,
	)
	mud.SetID(id)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		r.EXPECT().FindUserByEmailAndPass(gomock.Any(), gomock.Any()).Return(mud, nil)

		ud, err := srv.findUserByEmailAndPassService(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, mud, ud)
	})

	t.Run("when_user_not_found", func(t *testing.T) {
		r.EXPECT().FindUserByEmailAndPass(gomock.Any(), gomock.Any()).Return(
			nil,
			rest_err.NewNotFoundError("User not found."),
		)

		ud, err := srv.findUserByEmailAndPassService(email, password)

		assert.Nil(t, ud)
		assert.EqualValues(t, err.Code, 404)
	})

	t.Run("when_user_not_found", func(t *testing.T) {
		r.EXPECT().FindUserByEmailAndPass(email, password).Return(
			nil,
			rest_err.NewInternalServerError("Error trying to find user"),
		)

		ud, err := srv.findUserByEmailAndPassService(email, password)

		assert.Nil(t, ud)
		assert.EqualValues(t, err.Code, 500)
	})
}
