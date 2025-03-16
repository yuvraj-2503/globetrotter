package db

import (
	"context"
	"errors"
	"globetrotter/auth/models"
	"globetrotter/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDBImpl struct {
	collection *mongo.Collection
}

func NewUserDB(collection *mongo.Collection) *UserDBImpl {
	return &UserDBImpl{
		collection: collection,
	}
}

func (u *UserDBImpl) InsertUser(ctx *context.Context, user *models.User) error {
	_, err := u.collection.InsertOne(*ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return &common.AlreadyExistsError{Message: "User already exists"}
		}
		return err
	}
	return nil
}

func (u *UserDBImpl) GetByEmail(ctx *context.Context, email string) (*models.User, error) {
	result := u.collection.FindOne(*ctx, bson.M{"email": email})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, &common.NotFoundError{Message: "User not found"}
		}
		return nil, result.Err()
	}

	var user models.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
