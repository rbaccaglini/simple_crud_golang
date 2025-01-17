package mongodb

import (
	"context"
	"os"

	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	JOURNEY     = zap.String("journey", "DB Connection")
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	mongodb_url := os.Getenv(MONGODB_URL)
	mongodb_user_db := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_url))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Database cennected", JOURNEY)

	return client.Database(mongodb_user_db), nil
}
