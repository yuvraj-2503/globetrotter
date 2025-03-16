package db

import (
	"context"
	"globetrotter/invite/models"
)

type InviteDB interface {
	Insert(ctx *context.Context, invite *models.Invite) error
	GetByInvitee(ctx *context.Context, invitee string) (*models.Invite, error)
}
