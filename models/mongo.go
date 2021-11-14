package models

import "go.mongodb.org/mongo-driver/mongo"

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}
