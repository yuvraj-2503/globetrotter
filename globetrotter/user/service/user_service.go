package service

import (
	"context"
	"globetrotter/user/db"
)

type UserServiceImpl struct {
	userDB db.UserDB
}

func NewUserService(userDB db.UserDB) *UserServiceImpl {
	return &UserServiceImpl{userDB: userDB}
}

func (s *UserServiceImpl) RegisterUser(ctx *context.Context, userDetails *db.UserDetails) error {
	return s.userDB.InsertUser(ctx, userDetails)
}

func (s *UserServiceImpl) GetUserByUsername(ctx *context.Context, username string) (*db.UserDetails, error) {
	return s.userDB.GetUserByUsername(ctx, username)
}

func (s *UserServiceImpl) GetUserById(ctx *context.Context, userId string) (*db.UserDetails, error) {
	return s.userDB.GetUserById(ctx, userId)
}
