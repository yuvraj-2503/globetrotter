package service

import (
	"context"
	"globetrotter/user/db"
)

type UserService interface {
	RegisterUser(ctx *context.Context, userDetails *db.UserDetails) error
	GetUserByUsername(ctx *context.Context, username string) (*db.UserDetails, error)
	GetUserById(ctx *context.Context, userId string) (*db.UserDetails, error)
}

type ServerService interface {
	UpdateScore(ctx *context.Context, userId string, score int) error
}
