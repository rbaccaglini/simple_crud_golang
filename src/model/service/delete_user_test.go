package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := mocks.NewMockUserRepository(ctl)
	srv := NewUserDomainService(r)

	t.Run("when_delete_with_success", func(t *testing.T) {
		r.EXPECT().DeleteUser(gomock.Any()).Return(nil)
		err := srv.DeleteUserService("123")
		assert.Nil(t, err)
	})

	t.Run("when_delete_with_error", func(t *testing.T) {
		r.EXPECT().DeleteUser(gomock.Any()).Return(rest_err.NewInternalServerError("error"))
		err := srv.DeleteUserService("123")
		assert.NotNil(t, err)
	})
}
