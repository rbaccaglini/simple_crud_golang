package user_service

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService(t *testing.T) {

	ctlr := gomock.NewController(t)
	defer ctlr.Finish()

	r := mocks.NewMockUserRepository(ctlr)
	srv := NewUserDomainService(r)

	t.Run("FindAllUser::service_error", func(t *testing.T) {
		r.EXPECT().GetUsers().Return(nil, rest_err.NewInternalServerError("error"))

		uds, err := srv.FindAllUser()

		assert.NotNil(t, err)
		assert.Nil(t, uds)
	})
	t.Run("FindAllUser::success_empty_return", func(t *testing.T) {
		muds := []domain.UserDomainInterface{}

		r.EXPECT().GetUsers().Return(muds, nil)

		uds, err := srv.FindAllUser()

		assert.Nil(t, err)
		assert.Equal(t, 0, len(uds))
	})
	t.Run("FindAllUser::success", func(t *testing.T) {
		mud1 := domain.NewUserDomain(
			"test1@test.com",
			"123$%¨7",
			"Test1 Silva",
			21,
		)
		mud1.SetID(primitive.NewObjectID().String())

		mud2 := domain.NewUserDomain(
			"test2@test.com",
			"223$%¨7",
			"Test2 Silva",
			22,
		)
		mud2.SetID(primitive.NewObjectID().String())

		muds := []domain.UserDomainInterface{mud1, mud2}

		r.EXPECT().GetUsers().Return(muds, nil)

		uds, err := srv.FindAllUser()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(uds))
	})

	t.Run("FindUserById::service_error", func(t *testing.T) {

		r.EXPECT().GetUserById("123").Return(nil, rest_err.NewInternalServerError("error"))

		ud, err := srv.FindUserById("123")

		assert.NotNil(t, err)
		assert.Nil(t, ud)
	})
	t.Run("FindUserById::success", func(t *testing.T) {

		resp := domain.NewUserDomain(
			"test@test.com", "12#$56", "User Name", 21,
		)

		r.EXPECT().GetUserById("123").Return(resp, nil)

		ud, err := srv.FindUserById("123")

		assert.Nil(t, err)
		assert.NotNil(t, ud)
		assert.Equal(t, resp.GetEmail(), ud.GetEmail())
	})

	t.Run("FindUserByEmail::service_error", func(t *testing.T) {
		r.EXPECT().GetUserByEmail("test@test.com").Return(nil, rest_err.NewInternalServerError("error"))

		ud, err := srv.FindUserByEmail("test@test.com")

		assert.NotNil(t, err)
		assert.Nil(t, ud)
	})
	t.Run("FindUserByEmail::success", func(t *testing.T) {
		resp := domain.NewUserDomain(
			"test@test.com", "12#$56", "User Name", 21,
		)

		r.EXPECT().GetUserByEmail("test@test.com").Return(resp, nil)

		ud, err := srv.FindUserByEmail("test@test.com")

		assert.Nil(t, err)
		assert.NotNil(t, ud)
		assert.Equal(t, resp.GetEmail(), ud.GetEmail())
	})

	t.Run("CreateUser::email_already_registered_1", func(t *testing.T) {
		r.EXPECT().GetUserByEmail(gomock.Any()).
			Return(nil, rest_err.NewInternalServerError("error"))

		ud := domain.NewUserDomain(
			"test@test.com", "12#$56", "User Name", 21,
		)

		ud, err := srv.CreateUser(ud)

		assert.NotNil(t, err)
		assert.Nil(t, ud)
		assert.Contains(t, err.Error(), "email is already registered")
	})
	t.Run("CreateUser::email_already_registered_2", func(t *testing.T) {
		r.EXPECT().GetUserByEmail(gomock.Any()).
			Return(domain.NewUserDomain(
				"email", "password", "name", 21,
			),
				nil,
			)

		ud := domain.NewUserDomain(
			"test@test.com", "12#$56", "User Name", 21,
		)

		ud, err := srv.CreateUser(ud)

		assert.NotNil(t, err)
		assert.Nil(t, ud)
		assert.Contains(t, err.Error(), "email is already registered")
	})
	t.Run("CreateUser::service_error", func(t *testing.T) {

		r.EXPECT().GetUserByEmail(gomock.Any()).Return(nil, nil)
		r.EXPECT().InsertUser(gomock.Any()).Return(
			nil,
			rest_err.NewInternalServerError("repository_error"),
		)

		ud := domain.NewUserDomain(
			"test@test.com", "12#$56", "User Name", 21,
		)

		ud, err := srv.CreateUser(ud)

		assert.NotNil(t, err)
		assert.Nil(t, ud)
		assert.Equal(t, http.StatusInternalServerError, err.Code)
	})
	t.Run("CreateUser::success", func(t *testing.T) {
		uid := primitive.NewObjectID().Hex()
		repoResp := domain.NewUserDomain(
			"test@test.com", "12#$56", "User Name", 21,
		)
		repoResp.SetID(uid)

		r.EXPECT().GetUserByEmail(gomock.Any()).Return(nil, nil)
		r.EXPECT().InsertUser(gomock.Any()).Return(
			repoResp,
			nil,
		)

		ud := domain.NewUserDomain(
			"test@test.com", "12#$56", "User Name", 21,
		)

		ud, err := srv.CreateUser(ud)

		assert.Nil(t, err)
		assert.NotNil(t, ud)
		assert.Equal(t, uid, ud.GetID())
	})

	t.Run("DeleteUser::service_error", func(t *testing.T) {})
	t.Run("DeleteUser::success", func(t *testing.T) {})

	t.Run("UpdateUser::user_not_found", func(t *testing.T) {})
	t.Run("UpdateUser::no_update_to_do", func(t *testing.T) {})
	t.Run("UpdateUser::service_error", func(t *testing.T) {})
	t.Run("UpdateUser::success", func(t *testing.T) {})

	t.Run("Login::invalid_credentials", func(t *testing.T) {})
	t.Run("Login::token_generation_error", func(t *testing.T) {})
	t.Run("Login::success", func(t *testing.T) {})

}
