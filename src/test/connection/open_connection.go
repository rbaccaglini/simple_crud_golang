package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ory/dockertest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (database *mongo.Database, close func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
		return nil, nil
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	})
	if err != nil {
		log.Fatalf("Could not create mongo container: %s", err)
		return nil, nil
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://127.0.0.1:%s", resource.GetPort("27017/tcp"))))
	if err != nil {
		log.Println("Error trying to open connection")
		return nil, nil
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println("Error trying to open connection")
		return nil, nil
	}

	database = client.Database(os.Getenv("MONGODB_USER_DB"))
	close = func() {
		log.Println("Closing Docker resource")
		err := resource.Close()
		if err != nil {
			log.Println("Error closing Docker resource")
			return
		}
		log.Println("Docker resource closed successfully")
	}

	return database, close
}
