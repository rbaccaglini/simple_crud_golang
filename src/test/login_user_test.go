package tests

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/controller/model/request"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestLoginUser(t *testing.T) {

	t.Run("some incorrect fields", func(t *testing.T) {
		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		body := request.UserLoginRequest{
			Email:    "email_not_valid",
			Password: "123$%¨7",
		}
		b, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		// Act
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.Login(ctx)

		// Assert
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Email must be a valid email address")
	})

	t.Run("user not exists", func(t *testing.T) {
		// Preparing Database Records
		err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			Drop(context.Background())
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		body := request.UserLoginRequest{
			Email:    "test@test.com",
			Password: "123$%¨7",
		}
		b, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		// Act
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.Login(ctx)

		// Assert
		assert.EqualValues(t, http.StatusForbidden, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Invalid credentials")
	})

	t.Run("success", func(t *testing.T) {

		email := "test@test.com"
		password := "123$%¨7"
		id := primitive.NewObjectID()

		// Preparing Database Records
		d := model.NewUserDomain(email, password, "Test Silva", 20)
		d.EncryptPassword()
		ePass := d.GetPassword()

		_, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertOne(
				context.Background(),
				bson.M{
					"_id":      id,
					"email":    email,
					"password": ePass,
					"name":     "Test Silva",
					"age":      20,
				},
			)
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		body := request.UserLoginRequest{
			Email:    email,
			Password: password,
		}
		b, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		// Act
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.Login(ctx)

		log.Printf("Token: %s", recorder.Header().Get("Authorization"))

		// Assert
		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.True(t, VerifyToken(recorder.Header().Get("Authorization")))
	})
}
