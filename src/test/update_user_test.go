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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateUser(t *testing.T) {

	t.Run("invalid user id", func(t *testing.T) {
		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		body := request.UserUpdateRequest{
			Name: "Test Silva",
			Age:  20,
		}
		b, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		param := []gin.Param{
			{Key: "userId", Value: "123"},
		}

		// Act
		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
		UserController.UpdateUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Invalid user id")
	})

	t.Run("some incorrect fields", func(t *testing.T) {
		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		body := request.UserUpdateRequest{
			Name: "Test Silva",
			Age:  -1,
		}
		b, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		param := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}

		// Act
		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
		UserController.UpdateUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Some fields are invalid")
	})

	t.Run("user not exists", func(t *testing.T) {

		// Preparing Database Records
		_, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertOne(context.Background(), bson.M{"name": t.Name(), "email": "test@test.com", "age": 20})
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		body := request.UserUpdateRequest{
			Name: "Test Silva",
			Age:  22,
		}
		b, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		param := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}

		// Act
		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
		UserController.UpdateUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusNoContent, recorder.Code)
		assert.Empty(t, recorder.Body.String())
	})

	t.Run("success", func(t *testing.T) {

		// Preparing Database Records
		uid, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertOne(
				context.Background(),
				bson.M{
					"name":  t.Name(),
					"email": "test@test.com",
					"age":   20,
				},
			)
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		body := request.UserUpdateRequest{
			Name: "Test Silva",
			Age:  22,
		}
		b, _ := json.Marshal(body)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		param := []gin.Param{
			{Key: "userId", Value: uid.InsertedID.(primitive.ObjectID).Hex()},
		}

		// Act
		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
		UserController.UpdateUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusNoContent, recorder.Code)
		assert.Empty(t, recorder.Body.String())
	})

}
