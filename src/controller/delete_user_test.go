package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_DeleteUser(t *testing.T) {

	t.Run("invalid user id", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: "not_an_id"},
		}
		MakeRequest(ctx, p, url.Values{}, "DELETE", nil)

		ctr.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)
	})

	t.Run("error calling service", func(t *testing.T) {
		userId := primitive.NewObjectID().Hex()

		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: userId},
		}
		MakeRequest(ctx, p, url.Values{}, "DELETE", nil)
		s.EXPECT().DeleteUserService(userId).Return(rest_err.NewInternalServerError("error"))

		ctr.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)

	})

	t.Run("success", func(t *testing.T) {
		userId := primitive.NewObjectID().Hex()

		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: userId},
		}
		MakeRequest(ctx, p, url.Values{}, "DELETE", nil)

		s.EXPECT().DeleteUserService(gomock.Any()).Return(nil)
		ctr.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusNoContent, r.Code)
	})
}
