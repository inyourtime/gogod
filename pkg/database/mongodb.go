package database

import (
	"context"
	"gogod/config"
	"gogod/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MC *mongo.Client

func MongoDBConnect(cfg *config.Env) *mongo.Client {
	var err error
	// create connection pool
	MC, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Db.Mongo.Uri))
	if err != nil {
		logger.Error(err)
	}
	// Test ping to mongodb server
	if err = MC.Ping(context.TODO(), nil); err != nil {
		logger.Error(err)
	}
	logger.Info("Mongodb has been initialize")
	return MC
}

func GetCollection(cfg *config.Env, client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(cfg.Db.Mongo.Database).Collection(collectionName)
}
