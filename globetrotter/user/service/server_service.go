package service

import (
	"context"
	"globetrotter/user/db"
)

type ServerServiceImpl struct {
	userDB db.UserDB
}

func NewServerService(userDB db.UserDB) *ServerServiceImpl {
	return &ServerServiceImpl{userDB: userDB}
}

func (s *ServerServiceImpl) UpdateScore(ctx *context.Context, userId string, score int) error {
	return s.userDB.UpdateScore(ctx, userId, score)
}
