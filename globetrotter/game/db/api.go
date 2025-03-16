package db

import (
	"context"
	"globetrotter/game/models"
)

type DestinationDB interface {
	Insert(ctx *context.Context, destination *models.Destination) error
	InsertBulk(ctx *context.Context, destinations []*models.Destination) error
	GetRandomDestination(ctx *context.Context) (*models.Destination, error)
	GetDestinationByID(ctx *context.Context, id string) (*models.Destination, error)
}
