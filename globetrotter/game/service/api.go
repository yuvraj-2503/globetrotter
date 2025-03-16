package service

import (
	"context"
	"globetrotter/game/models"
)

type GameService interface {
	GetRandomQuestion(ctx *context.Context) (*models.Question, error)
	SubmitAnswer(ctx *context.Context, userId string, req *models.AnswerRequest) (*models.Verdict, error)
}

type DestinationService interface {
	Insert(ctx *context.Context, destination *models.Destination) error
	InsertBulk(ctx *context.Context, destinations []*models.Destination) error
}
