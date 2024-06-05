package model

import (
	"context"
	"github.com/run-bigpig/bragking/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var database *mongo.Database

func Init(ctx context.Context) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Get().Mongo.Url))
	if err != nil {
		log.Fatal(err)
	}
	database = client.Database(config.Get().Mongo.Database)
}

func GetDataBase() *mongo.Database {
	return database
}
