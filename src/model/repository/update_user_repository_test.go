package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/rbaccaglini/simple_crud_golang/src/model"
	entity "github.com/rbaccaglini/simple_crud_golang/src/model/repository/converter"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	userId := primitive.NewObjectID().Hex()
	dm := model.NewUserDomain(
		"test@test.com",
		"123$%6",
		"Test Silva",
		20,
	)

	mockEntity := entity.UserEntity{
		ID:       primitive.NewObjectID(),
		Email:    "test@test.com",
		Password: "test12%",
		Name:     "Test Silva",
		Age:      25,
	}

	mtestDb.Run("Test With Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(mockEntity),
		))
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		err := repo.UpdateUser(userId, dm)
		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		err := repo.UpdateUser(userId, dm)

		assert.NotNil(t, err)
	})
}
