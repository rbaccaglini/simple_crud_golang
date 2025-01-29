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

func TestUserDomainService_Login(t *testing.T) {
	t.Run("some invalid field", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserLoginRequest{
			Email:    "email_not_valid",
			Password: "123$%¨7",
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		ctr.Login(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "Some fields are invalid")
	})

	t.Run("error on calling service", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserLoginRequest{
			Email:    "test@test.com",
			Password: "123$%¨7",
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		s.EXPECT().
			LoginUserService(gomock.Any()).
			Return(nil, "", rest_err.NewInternalServerError("error"))

		ctr.Login(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("success", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserLoginRequest{
			Email:    "test@test.com",
			Password: "123$%¨7",
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		sud := model.NewUserDomain(
			"test@test.com",
			"pw12#$5",
			"Test Silva",
			20,
		)

		s.EXPECT().
			LoginUserService(gomock.Any()).
			Return(
				sud,
				"tk123456",
				nil)

		ctr.Login(ctx)

		ur := response.UserResponse{}
		body := r.Body.Bytes()
		err := json.Unmarshal(body, &ur)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.EqualValues(t, http.StatusOK, r.Code)
		assert.NotEmpty(t, r.Header().Get("Authorization"))
		assert.EqualValues(t, sud.GetEmail(), ur.Email)
		assert.EqualValues(t, sud.GetName(), ur.Name)
		assert.EqualValues(t, sud.GetAge(), ur.Age)
	})
}
