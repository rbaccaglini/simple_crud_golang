package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	request "github.com/rbaccaglini/simple_crud_golang/src/controller/model/request"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_UpdateUser(t *testing.T) {

	t.Run("Invalid user id", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: "not_an_id"},
		}

		MakeRequest(ctx, p, url.Values{}, "PUT", nil)
		ctr.UpdateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)
	})

	t.Run("Some invalid field", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserUpdateRequest{
			Name: "",
			Age:  -1,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		p := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}

		MakeRequest(ctx, p, url.Values{}, "PUT", stringReader)
		ctr.UpdateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "Some fields are invalid")
	})

	t.Run("Error calling service", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserUpdateRequest{
			Name: "",
			Age:  20,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		s.EXPECT().
			UpdateUserService(gomock.Any(), gomock.Any()).
			Return(rest_err.NewInternalServerError("error"))

		p := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}
		MakeRequest(ctx, p, url.Values{}, "PUT", stringReader)
		ctr.UpdateUser(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("success", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserUpdateRequest{
			Name: "",
			Age:  20,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		s.EXPECT().
			UpdateUserService(gomock.Any(), gomock.Any()).
			Return(nil)

		p := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}
		MakeRequest(ctx, p, url.Values{}, "PUT", stringReader)
		ctr.UpdateUser(ctx)

		assert.EqualValues(t, http.StatusNoContent, r.Code)
	})
}
