package db

import (
	"context"
	"errors"
	"globetrotter/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserStore struct {
	collection *mongo.Collection
}

func NewMongoUserStore(collection *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		collection: collection,
	}
}

func (m *MongoUserStore) InsertUser(ctx *context.Context, user *UserDetails) error {
	filter := bson.M{"userId": user.UserId}
	update := bson.M{"$set": user}
	opts := options.Update().SetUpsert(true)
	_, err := m.collection.UpdateOne(*ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUserStore) UpdateScore(ctx *context.Context, userId string, score int) error {
	filter := bson.M{"userId": userId}
	update := bson.M{"$inc": bson.M{"score": score}}

	_, err := m.collection.UpdateOne(*ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUserStore) GetUserByUsername(ctx *context.Context, username string) (*UserDetails, error) {
	result := m.collection.FindOne(*ctx, bson.M{"username": username})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, &common.NotFoundError{Message: "User not found"}
	}

	var userDetails UserDetails
	err := result.Decode(&userDetails)
	if err != nil {
		return nil, err
	}
	return &userDetails, nil
}

func (m *MongoUserStore) GetUserById(ctx *context.Context, userId string) (*UserDetails, error) {
	result := m.collection.FindOne(*ctx, bson.M{"userId": userId})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, &common.NotFoundError{Message: "User not found"}
		}
		return nil, result.Err()
	}
	var userDetails UserDetails
	err := result.Decode(&userDetails)
	if err != nil {
		return nil, err
	}
	return &userDetails, nil
}
