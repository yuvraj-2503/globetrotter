package service

import (
	"context"
	"globetrotter/game/db"
	"globetrotter/game/models"
)

type DestinationServiceImpl struct {
	destDB db.DestinationDB
}

func NewDestinationService(destDB db.DestinationDB) *DestinationServiceImpl {
	return &DestinationServiceImpl{destDB: destDB}
}

func (service *DestinationServiceImpl) Insert(ctx *context.Context, destination *models.Destination) error {
	return service.destDB.Insert(ctx, destination)
}

func (service *DestinationServiceImpl) InsertBulk(ctx *context.Context, destinations []*models.Destination) error {
	return service.destDB.InsertBulk(ctx, destinations)
}
