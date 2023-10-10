package database

import (
	"context"
	"gogod/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MC *mongo.Client

func MongoDBConnect(cfg *config.Env) *mongo.Client {
	var err error
	// create connection pool
	MC, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Db.Mongo.Uri))
	if err != nil {
		log.Fatalf("Connect to mongodb fail: %v", err)
	}
	// Test ping to mongodb server
	if err = MC.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Ping to mongodb error: %v", err)
	}
	log.Println("Mongodb has been initialize")
	return MC
}

func GetCollection(cfg *config.Env, client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(cfg.Db.Mongo.Database).Collection(collectionName)
}
