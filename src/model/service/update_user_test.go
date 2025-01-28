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

func TestUserDomainService_UpdateUserService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	r := mocks.NewMockUserRepository(ctl)
	srv := NewUserDomainService(r)

	t.Run("success", func(t *testing.T) {
		userId := primitive.NewObjectID().String()
		ud := model.NewUserDomain("test@test.com", "123$%¨7", "Test Silva", 20)

		r.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)

		err := srv.UpdateUserService(userId, ud)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		userId := primitive.NewObjectID().String()
		ud := model.NewUserDomain("test@test.com", "123$%¨7", "Test Silva", 20)

		r.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(rest_err.NewInternalServerError("error update"))

		err := srv.UpdateUserService(userId, ud)
		assert.NotNil(t, err)
		assert.EqualValues(t, 500, err.Code)
	})
}
