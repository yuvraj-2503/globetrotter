package db

import "context"

type UserDB interface {
	InsertUser(ctx *context.Context, user *UserDetails) error
	UpdateScore(ctx *context.Context, userId string, score int) error
	GetUserByUsername(ctx *context.Context, username string) (*UserDetails, error)
	GetUserById(ctx *context.Context, userId string) (*UserDetails, error)
}
