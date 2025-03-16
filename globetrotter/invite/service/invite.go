package service

import (
	"context"
	"globetrotter/config"
	"globetrotter/invite/db"
	"globetrotter/invite/models"
	"globetrotter/user/service"
)

type InviteServiceImpl struct {
	inviteDB    db.InviteDB
	userService service.UserService
}

func NewInviteService(inviteDB db.InviteDB, userService service.UserService) *InviteServiceImpl {
	return &InviteServiceImpl{
		inviteDB:    inviteDB,
		userService: userService,
	}
}

func (s *InviteServiceImpl) GetInviteLink(ctx *context.Context, userId string, invitee string) (string, error) {
	userdetails, err := s.userService.GetUserById(ctx, userId)
	if err != nil {
		return "", err
	}

	inviterUsername := userdetails.Username
	err = s.inviteDB.Insert(ctx, &models.Invite{Invitee: invitee, Inviter: inviterUsername})
	if err != nil {
		return "", err
	}

	inviteLink := config.Configuration.InviteBaseUrl + "?username=" + invitee
	return inviteLink, nil
}

func (s *InviteServiceImpl) GetInviterScore(ctx *context.Context, userId string) (*models.InviterScoreResponse, error) {
	userdetails, err := s.userService.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	invitee := userdetails.Username
	invite, err := s.inviteDB.GetByInvitee(ctx, invitee)
	if err != nil {
		return nil, err
	}

	inviter := invite.Inviter
	result, err := s.userService.GetUserByUsername(ctx, inviter)
	if err != nil {
		return nil, err
	}
	return &models.InviterScoreResponse{
		Inviter: result.Username,
		Score:   result.Score,
	}, nil
}
