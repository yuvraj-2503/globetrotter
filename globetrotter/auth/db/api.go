package db

import (
	"context"
	"globetrotter/auth/models"
)

type UserDB interface {
	InsertUser(ctx *context.Context, user *models.User) error
	GetByEmail(ctx *context.Context, email string) (*models.User, error)
}
