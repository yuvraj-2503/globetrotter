package game

import (
	"context"
	"globetrotter/common"
	"globetrotter/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	mongodb "mongo-utils"
)

func LoadDB(ctx *context.Context) {
	mongoConfig := config.Configuration.MongoConfig
	loadDestinationsCollection(ctx, mongoConfig)
}

func loadDestinationsCollection(ctx *context.Context, mongoConfig *mongodb.MongoConfig) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"city", 1},
			{"country", 1},
		},
		Options: options.Index().SetName("city-country-index").SetUnique(true),
	}

	userColl, err := mongoConfig.GetCollection(common.DestinationCollection)
	if err != nil {
		log.Fatalf("Failed to connect to database, reason: %v", err.Error())
	}

	indexes, err := userColl.Indexes().CreateMany(*ctx, []mongo.IndexModel{indexModel})
	if err != nil {
		log.Fatalf("Failed to create indexes, reason: %v", err.Error())
	}
	log.Printf("Successfully created indexes with name: %v", indexes)
}
