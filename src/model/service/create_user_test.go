package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserDomainService_CreateUserService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := mocks.NewMockUserRepository(ctl)
	srv := NewUserDomainService(r)

	t.Run("when_user_not_existis_returns_success", func(t *testing.T) {

		repoResp := model.NewUserDomain(
			"test@test.com", "123$%6", "Test Silva", 21,
		)

		r.EXPECT().FindUserByEmail(gomock.Any()).Return(nil, rest_err.NewNotFoundError("user not found"))
		r.EXPECT().CreateUser(gomock.Any()).Return(repoResp, nil)

		input := model.NewUserDomain(
			"test@test.com", "123$%6", "Test Silva", 21,
		)
		r, err := srv.CreateUserService(input)

		assert.Nil(t, err)
		assert.EqualValues(t, repoResp, r)
	})

	t.Run("when_email_already_registered", func(t *testing.T) {

		repoResp := model.NewUserDomain(
			"test@test.com", "123$%6", "Test Silva", 21,
		)

		r.EXPECT().FindUserByEmail(gomock.Any()).
			Return(repoResp, nil)

		input := model.NewUserDomain(
			"test@test.com", "123$%6", "Test Silva", 21,
		)
		r, err := srv.CreateUserService(input)

		assert.Nil(t, r)
		assert.EqualValues(t, err.Message, "Email is already registered")
		assert.EqualValues(t, err.Code, 400)
	})

	t.Run("when_error_find_user", func(t *testing.T) {

		r.EXPECT().FindUserByEmail(gomock.Any()).
			Return(nil, rest_err.NewBadRequestError(""))

		input := model.NewUserDomain(
			"test@test.com", "123$%6", "Test Silva", 21,
		)
		r, err := srv.CreateUserService(input)

		assert.Nil(t, r)
		assert.EqualValues(t, err.Message, "Email is already registered")
		assert.EqualValues(t, err.Code, 400)
	})

	t.Run("when_error_on_create_repo", func(t *testing.T) {
		r.EXPECT().FindUserByEmail(gomock.Any()).Return(nil, rest_err.NewNotFoundError("user not found"))
		r.EXPECT().CreateUser(gomock.Any()).Return(nil, rest_err.NewInternalServerError("error on repository"))

		input := model.NewUserDomain(
			"test@test.com", "123$%6", "Test Silva", 21,
		)
		r, err := srv.CreateUserService(input)

		assert.Nil(t, r)
		assert.EqualValues(t, err.Message, "error on repository")
	})
}
