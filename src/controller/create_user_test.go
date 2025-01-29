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
	"github.com/rbaccaglini/simple_crud_golang/src/controller/model/response"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserDomainService_CreateUser(t *testing.T) {

	t.Run("Some incorrect fields", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserRequest{
			Email:    "email_not_valid",
			Password: "123$%¨7",
			Name:     "Test User",
			Age:      0,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		ctr.CreateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "Some fields are invalid")

	})

	t.Run("Error calling service", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		s.EXPECT().
			CreateUserService(gomock.Any()).
			Return(
				nil,
				rest_err.NewInternalServerError("error"),
			)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "123$%¨7",
			Name:     "Test User",
			Age:      20,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		ctr.CreateUser(ctx)
		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("success", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		s.EXPECT().
			CreateUserService(gomock.Any()).
			Return(
				model.NewUserDomain(
					"test@test.com",
					"123$%¨7",
					"Test User",
					20),
				nil,
			)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "123$%¨7",
			Name:     "Test User",
			Age:      20,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		ctr.CreateUser(ctx)

		ur := response.UserResponse{}
		br := r.Body.Bytes()
		err := json.Unmarshal(br, &ur)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.EqualValues(t, http.StatusCreated, r.Code)
		assert.EqualValues(t, userRequest.Age, ur.Age)
		assert.EqualValues(t, userRequest.Email, ur.Email)
		assert.EqualValues(t, userRequest.Name, ur.Name)
	})
}
