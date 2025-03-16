package service

import (
	"context"
	"globetrotter/invite/models"
)

type InviteService interface {
	GetInviteLink(ctx *context.Context, userId string, invitee string) (string, error)
	GetInviterScore(ctx *context.Context, userId string) (*models.InviterScoreResponse, error)
}
