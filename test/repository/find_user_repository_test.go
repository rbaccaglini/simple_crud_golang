package user_repository_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rbaccaglini/simple_crud_golang/internal/models/entity"
	user_repository "github.com/rbaccaglini/simple_crud_golang/internal/repositories/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type Config struct {
	Port             string
	MongoURI         string
	DatabaseName     string
	UserDbCollection string
	JWTSecret        string
}

func TestUserRepository_GetUserById(t *testing.T) {
	databaseName := "user_database_test"
	config := &Config{
		Port:             "8080",
		MongoURI:         "mongodb://localhost:27017",
		DatabaseName:     "crud-init",
		UserDbCollection: "users",
		JWTSecret:        "123456",
	}
	uid := primitive.NewObjectID().Hex()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("GetUserById::mongodb_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)
		repo := user_repository.NewUserRepository(databaseMock)

		userDomain, err := repo.GetUserById("test")

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Error trying to find user")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("GetUserById::user_not_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, config.UserDbCollection),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := user_repository.NewUserRepository(databaseMock)
		userDomain, err := repo.GetUserById("test")

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "User not found.")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("GetUserById::success", func(mt *mtest.T) {
		mockEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test12%",
			Name:     "Test Silva",
			Age:      25,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, config.UserDbCollection),
			mtest.FirstBatch,
			convertEntityToBson(mockEntity),
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := user_repository.NewUserRepository(databaseMock)

		userDomain, err := repo.GetUserById(uid)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), mockEntity.Email)
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	email := "test@test.com"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("GetUserById::mongodb_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := user_repository.NewUserRepository(databaseMock)
		userDomain, err := repo.GetUserByEmail(email)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Error trying to find user")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("GetUserById::user_not_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := user_repository.NewUserRepository(databaseMock)
		userDomain, err := repo.GetUserByEmail(email)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "User not found.")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("GetUserById::success", func(mt *mtest.T) {
		mockEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test12%",
			Name:     "Test Silva",
			Age:      25,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(mockEntity),
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := user_repository.NewUserRepository(databaseMock)

		userDomain, err := repo.GetUserByEmail(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), mockEntity.Email)
	})
}

func TestUserRepository_GetUsers(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("GetUsers::mongodb_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := user_repository.NewUserRepository(databaseMock)
		userDomain, err := repo.GetUsers()

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Error trying to get all users")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("GetUsers::empty_return", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := user_repository.NewUserRepository(databaseMock)
		userDomains, err := repo.GetUsers()

		assert.Nil(t, err)
		assert.EqualValues(t, len(userDomains), 0)
	})

	mtestDb.Run("GetUserById::success", func(mt *mtest.T) {
		mockEntity1 := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test12%",
			Name:     "Test Silva",
			Age:      25,
		}

		mockEntity2 := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test12%",
			Name:     "Test Silva",
			Age:      25,
		}

		mockEntities := []entity.UserEntity{mockEntity1, mockEntity2}
		bsonEntities := []bson.D{}
		for _, e := range mockEntities {
			bsonEntities = append(bsonEntities, convertEntityToBson(e))
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			int64(len(mockEntities)),
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			bsonEntities...,
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := user_repository.NewUserRepository(databaseMock)
		userDomains, err := repo.GetUsers()

		assert.Nil(t, err)
		assert.EqualValues(t, len(userDomains), 2)
		assert.EqualValues(t, mockEntities[0].Email, userDomains[0].GetEmail())
	})
}

func TestUserRepository_ValidateCredentials(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	email := "test@test.com"
	password := "test123$"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("GetUserById::mongodb_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := user_repository.NewUserRepository(databaseMock)
		userDomain, err := repo.ValidateCredentials(email, password)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Error trying to find user")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("GetUserById::user_not_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := user_repository.NewUserRepository(databaseMock)
		userDomain, err := repo.ValidateCredentials(email, password)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "User not found.")
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("GetUserById::success", func(mt *mtest.T) {
		mockEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test12%",
			Name:     "Test Silva",
			Age:      25,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(mockEntity),
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := user_repository.NewUserRepository(databaseMock)

		userDomain, err := repo.ValidateCredentials(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), mockEntity.Email)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
