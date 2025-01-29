package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/controller/model/response"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/rbaccaglini/simple_crud_golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_FindUserById(t *testing.T) {
	mCtr := gomock.NewController(t)
	s := mocks.NewMockUserDomainService(mCtr)
	ctr := NewUserControllerInterface(s)

	userId := primitive.NewObjectID().Hex()

	t.Run("invalid id", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)
		p := []gin.Param{
			{Key: "userId", Value: "not_an_id"},
		}

		MakeRequest(ctx, p, url.Values{}, "GET", nil)
		ctr.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)
	})

	t.Run("Error on call service", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)
		p := []gin.Param{
			{Key: "userId", Value: userId},
		}

		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		s.EXPECT().FindUserByIdService(gomock.Any()).Return(
			nil,
			rest_err.NewInternalServerError("error"),
		)

		ctr.FindUserById(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("success", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)
		p := []gin.Param{
			{Key: "userId", Value: userId},
		}

		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		s.EXPECT().FindUserByIdService(gomock.Any()).Return(
			model.NewUserDomain("test@test.com", "", "Test Silva", 20),
			nil,
		)

		ctr.FindUserById(ctx)
		assert.EqualValues(t, http.StatusOK, r.Code)
	})

}

func TestUserDomainService_FindUserByEmail(t *testing.T) {

	mCtr := gomock.NewController(t)
	s := mocks.NewMockUserDomainService(mCtr)
	ctr := NewUserControllerInterface(s)

	t.Run("invalid email", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)
		p := []gin.Param{
			{Key: "userEmail", Value: "not_a_email"},
		}

		MakeRequest(ctx, p, url.Values{}, "GET", nil)
		ctr.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusBadRequest, r.Code)

	})

	t.Run("Error on call service", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)
		p := []gin.Param{
			{Key: "userEmail", Value: "test@test.com"},
		}

		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		s.EXPECT().FindUserByEmailService(gomock.Any()).Return(
			nil,
			rest_err.NewInternalServerError("error"),
		)

		ctr.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)

	})

	t.Run("success", func(t *testing.T) {
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)
		p := []gin.Param{
			{Key: "userEmail", Value: "test@test.com"},
		}

		MakeRequest(ctx, p, url.Values{}, "GET", nil)

		s.EXPECT().FindUserByEmailService(gomock.Any()).Return(
			model.NewUserDomain("test@test.com", "", "Test Silva", 20),
			nil,
		)

		ctr.FindUserByEmail(ctx)
		assert.EqualValues(t, http.StatusOK, r.Code)
	})
}

func TestUserDomainService_FindAllUsers(t *testing.T) {
	t.Run("Error on call service", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		MakeRequest(ctx, nil, url.Values{}, "GET", nil)
		s.EXPECT().FindAllUsersService().Return(nil, rest_err.NewInternalServerError("error"))

		ctr.FindAllUsers(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("success", func(t *testing.T) {
		mCtr := gomock.NewController(t)
		s := mocks.NewMockUserDomainService(mCtr)
		ctr := NewUserControllerInterface(s)

		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		srvRet := []model.UserDomainInterface{
			model.NewUserDomain("test1@test.com", "", "Test1 Silva", 20),
			model.NewUserDomain("test2@test.com", "", "Test2 Silva", 21),
		}

		MakeRequest(ctx, nil, url.Values{}, "GET", nil)
		s.EXPECT().FindAllUsersService().Return(
			srvRet,
			nil,
		)

		ctr.FindAllUsers(ctx)

		ur := []response.UserResponse{}
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
