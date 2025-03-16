package db

import (
	"context"
	"errors"
	"globetrotter/invite/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"user-server/common"
)

type InviteDBImpl struct {
	collection *mongo.Collection
}

func NewInviteDB(collection *mongo.Collection) *InviteDBImpl {
	return &InviteDBImpl{collection: collection}
}

func (i *InviteDBImpl) Insert(ctx *context.Context, invite *models.Invite) error {
	_, err := i.collection.InsertOne(*ctx, invite)
	if err != nil {
		return err
	}
	return nil
}

func (i *InviteDBImpl) GetByInvitee(ctx *context.Context, invitee string) (*models.Invite, error) {
	filter := bson.M{"invitee": invitee}
	result := i.collection.FindOne(*ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, &common.NotFoundError{Message: "Invite not found"}
		}
		return nil, result.Err()
	}
	var invite models.Invite
	err := result.Decode(&invite)
	if err != nil {
		return nil, err
	}
	return &invite, nil
}
