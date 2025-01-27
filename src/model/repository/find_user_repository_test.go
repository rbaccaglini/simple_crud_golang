package repository

import (
	"fmt"
	"os"
	"testing"

	entity "github.com/rbaccaglini/simple_crud_golang/src/model/repository/converter"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("Test With Success", func(mt *mtest.T) {

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
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmail("test@test.com")

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), mockEntity.Email)

	})

	mtestDb.Run("Test With Mongo Error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("Test With No Document Found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserById(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	userId := primitive.NewObjectID()

	mtestDb.Run("Test With Success", func(mt *mtest.T) {
		mockEntity := entity.UserEntity{
			ID:       userId,
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
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserById(userId.String())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), mockEntity.Email)

	})

	mtestDb.Run("Test With Mongo Error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById(userId.String())

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("Test With No Document Found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById(userId.String())

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindAllUsers(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("Test Find All With Success", func(mt *mtest.T) {

		userId1 := primitive.NewObjectID()
		userId2 := primitive.NewObjectID()

		mockEntities := []entity.UserEntity{
			{
				ID:       userId1,
				Email:    "test1@test.com",
				Password: "test12%",
				Name:     "Test Silva 1",
				Age:      25,
			},
			{
				ID:       userId2,
				Email:    "test2@test.com",
				Password: "test12%",
				Name:     "Test Silva 2",
				Age:      30,
			},
		}

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
		repo := NewUserRepository(databaseMock)

		userDomains, err := repo.FindAllUsers()

		assert.Nil(t, err)
		assert.EqualValues(t, len(userDomains), 2)
		assert.EqualValues(t, mockEntities[0].Email, userDomains[0].GetEmail())
	})

	mtestDb.Run("Test With Mongo Error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindAllUsers()

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("Test With No Document Found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindAllUsers()

		assert.Nil(t, err)
		assert.EqualValues(t, len(userDomain), 0)
	})
}

func TestUserRepository_FindUserByEmailAndPass(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	email := "test@test.com"
	password := "123$%62"
	userId := primitive.NewObjectID()

	mtestDb.Run("Test With Success", func(mt *mtest.T) {

		mockEntity := entity.UserEntity{
			ID:       userId,
			Email:    email,
			Password: password,
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
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPass(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), mockEntity.Email)
	})

	mtestDb.Run("Test With Mongo Error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPass(email, password)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("Test With No Document Found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPass(email, password)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
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
