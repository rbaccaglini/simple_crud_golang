package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/controller/model/request"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateUser(t *testing.T) {

	t.Run("some_incorrect_fields", func(t *testing.T) {
		// Arrange
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

		// Act
		UserController.CreateUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "Email must be a valid email address")

	})

	t.Run("user_already_registered", func(t *testing.T) {

		// Preparing Database Records
		_, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertOne(context.Background(), bson.M{"name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "123$%¨7",
			Name:     "Test User",
			Age:      20,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		// Act
		UserController.CreateUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "Email is already registered")
	})

	t.Run("user_create_success", func(t *testing.T) {
		// Preparing Database Records
		err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			Drop(context.Background())
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "123$%¨7",
			Name:     "Test User",
			Age:      20,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		// Act
		UserController.CreateUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusCreated, r.Code)
		assert.Contains(t, r.Body.String(), "\"email\":\"test@test.com\",\"name\":\"Test User\",\"age\":20}")
	})
}
