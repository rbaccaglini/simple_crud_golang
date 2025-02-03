package user_handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	user_handler "github.com/rbaccaglini/simple_crud_golang/internal/handlers/user"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	user_response "github.com/rbaccaglini/simple_crud_golang/internal/models/response/user"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_FindAllUser(t *testing.T) {

	mCtr := gomock.NewController(t)
	s := mocks.NewMockUserDomainService(mCtr)
	ctr := user_handler.NewUserControllerInterface(s)

	t.Run("FindAllUser::empty_response", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		srvRet := []domain.UserDomainInterface{}

		MakeRequest(ctx, nil, url.Values{}, "GET", nil)
		s.EXPECT().FindAllUser().Return(
			srvRet,
			nil,
		)

		ctr.FindAllUser(ctx)

		ur := []user_response.UserResponse{}
		b := r.Body.String()
		err := json.Unmarshal([]byte(b), &ur)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.EqualValues(t, http.StatusOK, r.Code)
		assert.EqualValues(t, 0, len(ur))
	})

	t.Run("FindAllUser::service_error", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		MakeRequest(ctx, nil, url.Values{}, "GET", nil)
		s.EXPECT().FindAllUser().Return(
			nil,
			rest_err.NewInternalServerError("error"),
		)

		ctr.FindAllUser(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("FindAllUser::success", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		srvRet := []domain.UserDomainInterface{
			domain.NewUserDomain("test1@test.com", "", "Test1 Silva", 20),
			domain.NewUserDomain("test2@test.com", "", "Test2 Silva", 21),
		}

		MakeRequest(ctx, nil, url.Values{}, "GET", nil)
		s.EXPECT().FindAllUser().Return(
			srvRet,
			nil,
		)

		ctr.FindAllUser(ctx)

		ur := []user_response.UserResponse{}
		b := r.Body.String()
		err := json.Unmarshal([]byte(b), &ur)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.EqualValues(t, http.StatusOK, r.Code)
		assert.EqualValues(t, len(srvRet), len(ur))
		assert.EqualValues(t, srvRet[0].GetEmail(), ur[0].Email)
	})
}

func TestUserDomainService_FindUserById(t *testing.T) {

	mCtr := gomock.NewController(t)
	s := mocks.NewMockUserDomainService(mCtr)
	ctr := user_handler.NewUserControllerInterface(s)

	t.Run("FindUserById::invalid_user_id", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: "invalid_user_id"},
		}

		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		ctr.FindUserById(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "Invalid user id")
	})

	t.Run("FindUserById::service_error", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}
		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		s.EXPECT().FindUserById(gomock.Any()).Return(
			nil, rest_err.NewInternalServerError("error_service"),
		)

		ctr.FindUserById(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("FindUserById::success", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		srvRet := domain.NewUserDomain("test1@test.com", "", "Test1 Silva", 20)

		p := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}
		MakeRequest(ctx, p, url.Values{}, "GET", nil)
		s.EXPECT().FindUserById(gomock.Any()).Return(
			srvRet,
			nil,
		)

		ctr.FindUserById(ctx)

		ur := user_response.UserResponse{}
		b := r.Body.String()
		err := json.Unmarshal([]byte(b), &ur)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.EqualValues(t, http.StatusOK, r.Code)
		assert.NotNil(t, ur)
		assert.EqualValues(t, srvRet.GetEmail(), ur.Email)
	})
}

func TestUserDomainService_FindUserByEmail(t *testing.T) {

	mCtr := gomock.NewController(t)
	s := mocks.NewMockUserDomainService(mCtr)
	ctr := user_handler.NewUserControllerInterface(s)

	t.Run("FindUserByEmail::invalid_user_email", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "email", Value: "invalid_user_id"},
		}
		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		ctr.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusNotFound, r.Code)
		assert.Contains(t, r.Body.String(), "user not found")
	})

	t.Run("FindUserByEmail::service_error", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "email", Value: "test@test.com"},
		}
		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		s.EXPECT().FindUserByEmail(gomock.Any()).Return(
			nil, rest_err.NewInternalServerError("error_service"),
		)

		ctr.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("FindUserByEmail::success", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		srvRet := domain.NewUserDomain("test1@test.com", "", "Test1 Silva", 20)

		p := []gin.Param{
			{Key: "email", Value: "test1@test.com"},
		}
		MakeRequest(ctx, p, url.Values{}, "GET", nil)
		s.EXPECT().FindUserByEmail(gomock.Any()).Return(
			srvRet,
			nil,
		)

		ctr.FindUserByEmail(ctx)

		ur := user_response.UserResponse{}
		b := r.Body.String()
		err := json.Unmarshal([]byte(b), &ur)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.EqualValues(t, http.StatusOK, r.Code)
		assert.NotNil(t, ur)
		assert.EqualValues(t, srvRet.GetEmail(), ur.Email)
	})
}

func GetTestGinContext(r *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(r)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func MakeRequest(c *gin.Context, param gin.Params, u url.Values, m string, b io.ReadCloser) {
	c.Request.Method = m
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.Body = b
	c.Request.URL.RawQuery = u.Encode()
}
