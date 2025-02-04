package integration_tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	user_response "github.com/rbaccaglini/simple_crud_golang/internal/models/response/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestIntegrationFindUserByEmail(t *testing.T) {

	t.Run("user_not_found_with_this_email", func(t *testing.T) {
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

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		// Act
		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserHandler.FindUserByEmail(ctx)

		// Assert
		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_with_specified_email", func(t *testing.T) {
		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		param := []gin.Param{
			{
				Key:   "email",
				Value: "test@test.com",
			},
		}

		// Preparing Database Records
		_, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertOne(context.Background(), bson.M{"name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		// Act
		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserHandler.FindUserByEmail(ctx)

		// Assert
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestFindUserById(t *testing.T) {

	t.Run("user_not_found_with_this_id", func(t *testing.T) {
		// Arrange
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		// Act
		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserHandler.FindUserById(ctx)

		// Assert
		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_with_specified_id", func(t *testing.T) {
		// Arrange
		id := primitive.NewObjectID()
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}
		MakeRequest(ctx, param, url.Values{}, "GET", nil)

		// Preparing Database Records
		_, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		// Act
		UserHandler.FindUserById(ctx)

		// Assert
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestFindAllUsers(t *testing.T) {
	t.Run("empty user list", func(t *testing.T) {
		// Arrange
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		// Drop all data in DB
		if err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			Drop(context.Background()); err != nil {
			t.Fatal(err)
			return
		}

		// Act
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "GET", nil)
		UserHandler.FindAllUser(ctx)

		ur := []user_response.UserResponse{}
		b := r.Body.String()
		err := json.Unmarshal([]byte(b), &ur)
		if err != nil {
			t.Fatal(err.Error())
		}

		// Assert
		assert.EqualValues(t, http.StatusOK, r.Code)
		assert.EqualValues(t, 0, len(ur))
	})

	t.Run("user list", func(t *testing.T) {

		// Preparing Database Records
		_, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertMany(
				context.Background(),
				[]interface{}{
					bson.M{"name": t.Name(), "email": "test1@test.com"},
					bson.M{"name": t.Name(), "email": "test2@test.com"},
				},
			)
		if err != nil {
			t.Fatal(err)
			return
		}

		count, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			CountDocuments(context.Background(), bson.M{})
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		// Act
		MakeRequest(ctx, []gin.Param{}, url.Values{}, "GET", nil)
		UserHandler.FindAllUser(ctx)

		ur := []user_response.UserResponse{}
		b := r.Body.Bytes()
		err = json.Unmarshal(b, &ur)
		if err != nil {
			t.Fatal(err.Error())
			return
		}

		// Assert
		assert.EqualValues(t, http.StatusOK, r.Code)
		assert.EqualValues(t, count, len(ur))
	})
}
