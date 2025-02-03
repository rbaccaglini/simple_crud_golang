package user_repository

import (
	"net/http"
	"testing"

	"github.com/rbaccaglini/simple_crud_golang/config"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_InsertUser(t *testing.T) {
	databaseName := "user_database_test"
	config := &config.Config{
		Port:             "8080",
		MongoURI:         "mongodb://localhost:27017",
		DatabaseName:     "crud-init",
		UserDbCollection: "users",
		JWTSecret:        "123456",
	}

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("InsertUser::success", func(mt *mtest.T) {
		mt.AddMockResponses(
			bson.D{{Key: "ok", Value: 1},
				{Key: "n", Value: 1},
				{Key: "acknowledged", Value: true},
			},
		)

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)
		domain := domain.NewUserDomain(
			"test@test.com", "test", "test", 90)
		userDomain, err := repo.InsertUser(domain)

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), domain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), domain.GetAge())
		assert.EqualValues(t, userDomain.GetPassword(), domain.GetPassword())

	})

	mtestDb.Run("InsertUser::error_mongodb", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)
		domain := domain.NewUserDomain(
			"test@test.com", "test", "test", 90)
		userDomain, err := repo.InsertUser(domain)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.Code)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_DeleteUser(t *testing.T) {
	databaseName := "user_database_test"
	config := &config.Config{
		Port:             "8080",
		MongoURI:         "mongodb://localhost:27017",
		DatabaseName:     "crud-init",
		UserDbCollection: "users",
		JWTSecret:        "123456",
	}

	uid := primitive.NewObjectID().Hex()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("DeleteUser::success", func(mt *mtest.T) {
		mt.AddMockResponses(
			bson.D{{Key: "ok", Value: 1},
				{Key: "n", Value: 1},
				{Key: "acknowledged", Value: true},
			},
		)

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)
		err := repo.DeleteUser(uid)

		assert.Nil(t, err)
	})

	mtestDb.Run("DeleteUser::invalid_user_id", func(mt *mtest.T) {
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)
		err := repo.DeleteUser("invalid_id")

		assert.NotNil(t, err)
		assert.Contains(t, err.Message, "invalid user id")
	})

	mtestDb.Run("DeleteUser::mongodb_error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)
		err := repo.DeleteUser(uid)

		assert.NotNil(t, err)
		assert.Contains(t, err.Message, "error on delete user")
	})
}

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	config := &config.Config{
		Port:             "8080",
		MongoURI:         "mongodb://localhost:27017",
		DatabaseName:     "crud-init",
		UserDbCollection: "users",
		JWTSecret:        "123456",
	}

	uid := primitive.NilObjectID.Hex()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mtestDb.Run("UpdateUser::invalid_user_id", func(mt *mtest.T) {
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)

		ud := domain.NewLoginDomain("test@test.com", "test123")

		err := repo.UpdateUser(ud, "invalid_id")

		assert.NotNil(t, err)
		assert.Contains(t, err.Message, "invalid user id")
	})

	mtestDb.Run("UpdateUser::mongodb_error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)

		ud := domain.NewUserUpdateDomain(uid, "User Name", 20)

		err := repo.UpdateUser(ud, uid)

		assert.NotNil(t, err)
		assert.Contains(t, err.Message, "error on update user with id")
	})

	mtestDb.Run("UpdateUser::success", func(mt *mtest.T) {

		mt.AddMockResponses(
			bson.D{{Key: "ok", Value: 1},
				{Key: "n", Value: 1},
				{Key: "acknowledged", Value: true},
			},
		)

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock, config)

		ud := domain.NewUserUpdateDomain(uid, "User Name", 20)

		err := repo.UpdateUser(ud, uid)

		assert.Nil(t, err)
	})

}
