package db

import (
	"context"
	"errors"
	"globetrotter/game/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"user-server/common"
)

type DestinationDBStore struct {
	collection *mongo.Collection
}

func NewDestinationDBStore(collection *mongo.Collection) *DestinationDBStore {
	return &DestinationDBStore{
		collection: collection,
	}
}

func (store *DestinationDBStore) Insert(ctx *context.Context, destination *models.Destination) error {
	_, err := store.collection.InsertOne(*ctx, destination)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return &common.AlreadyExistsError{Message: "Destination already exists"}
		}
		return err
	}
	return nil
}

func (store *DestinationDBStore) InsertBulk(ctx *context.Context, destinations []*models.Destination) error {
	if len(destinations) == 0 {
		return nil
	}

	var operations []mongo.WriteModel

	for _, dest := range destinations {
		filter := bson.M{"city": dest.City, "country": dest.Country} // Unique identifier
		update := bson.M{
			"$set": bson.M{
				"clues":     dest.Clues,
				"fun_facts": dest.FunFacts,
				"trivia":    dest.Trivia,
				"options":   dest.Options,
			},
		}

		updateModel := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		operations = append(operations, updateModel)
	}

	_, err := store.collection.BulkWrite(*ctx, operations)
	if err != nil {
		return err
	}
	return nil
}

func (store *DestinationDBStore) GetRandomDestination(ctx *context.Context) (*models.Destination, error) {
	count, err := store.collection.CountDocuments(*ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, &common.NotFoundError{Message: "No destinations found."}
	}

	skip := rand.Int63n(count)
	limit := int64(1)
	cursor, err := store.collection.Find(*ctx, bson.M{}, &options.FindOptions{Skip: &skip, Limit: &limit})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(*ctx)

	var destination models.Destination
	if cursor.Next(*ctx) {
		if err := cursor.Decode(&destination); err != nil {
			return nil, err
		}
		return &destination, nil
	}
	return nil, &common.NotFoundError{Message: "No destinations found."}
}

func (store *DestinationDBStore) GetDestinationByID(ctx *context.Context, id string) (*models.Destination, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	result := store.collection.FindOne(*ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, &common.NotFoundError{Message: "No destination found for id: " + id}
		}
		return nil, result.Err()
	}
	var destination models.Destination
	if err := result.Decode(&destination); err != nil {
		return nil, err
	}
	return &destination, nil
}
