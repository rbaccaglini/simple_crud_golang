package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteUser(t *testing.T) {

	t.Run("Invalid user id", func(t *testing.T) {
		// Arrange
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: "not_an_id"},
		}
		MakeRequest(ctx, p, url.Values{}, "DELETE", nil)

		// Act
		UserController.DeleteUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "Invalid user id")

	})

	t.Run("user_not_exists", func(t *testing.T) {

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

		p := []gin.Param{
			{Key: "userId", Value: primitive.NewObjectID().Hex()},
		}
		MakeRequest(ctx, p, url.Values{}, "DELETE", nil)

		// Act
		UserController.DeleteUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusNoContent, r.Code)
		assert.Empty(t, r.Body.String())
	})

	t.Run("user_create_success", func(t *testing.T) {
		// Preparing Database Records
		uid, err := Database.
			Collection(os.Getenv("MONGODB_USER_DB_COLLECTION")).
			InsertOne(context.Background(), bson.M{"name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		// Arrange
		r := httptest.NewRecorder()
		ctx := GetTestGinContext(r)

		p := []gin.Param{
			{Key: "userId", Value: uid.InsertedID.(primitive.ObjectID).Hex()},
		}
		MakeRequest(ctx, p, url.Values{}, "DELETE", nil)

		// Act
		UserController.DeleteUser(ctx)

		// Assert
		assert.EqualValues(t, http.StatusNoContent, r.Code)
		assert.Empty(t, r.Body.String())
	})
}
