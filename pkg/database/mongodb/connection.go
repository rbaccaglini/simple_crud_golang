package mongodb

import (
	"context"

	"github.com/rbaccaglini/simple_crud_golang/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	JOURNEY     = zap.String("journey", "DB Connection")
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context, cfg *config.Config) (*mongo.Database, error) {

	mongodb_url := cfg.MongoURI
	mongodb_user_db := cfg.DatabaseName

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_url))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_user_db), nil
}
