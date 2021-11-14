package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/farukak/go_test_api/config"
	"github.com/farukak/go_test_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mg models.MongoInstance

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
func Connect() error {
	uri := config.Config("MONGO_URI")
	dbName := config.Config("DB_NAME")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = models.MongoInstance{
		Client: client,
		Db:     db,
	}

	//defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return err
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return err
	}

	fmt.Println(databases)

	return nil
}
