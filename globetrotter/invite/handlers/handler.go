package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	common2 "globetrotter/common"
	"globetrotter/invite/service"
	"net/http"
	token "token-manager"
)

type InviteHandler struct {
	service service.InviteService
}

func NewInviteHandler(service service.InviteService) *InviteHandler {
	return &InviteHandler{service: service}
}

func (i *InviteHandler) GetInviteLink(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	invitee := ctx.Query("invitee")
	userid := getUserIdFromContext(ctx)

	inviteLink, err := i.service.GetInviteLink(&requestCtx, userid, invitee)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"inviteLink": inviteLink,
	})
}

func (i *InviteHandler) GetInviterScore(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	userid := getUserIdFromContext(ctx)

	result, err := i.service.GetInviterScore(&requestCtx, userid)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func getUserIdFromContext(ctx *gin.Context) string {
	user, exists := ctx.Get("user")
	if !exists {
		return ""
	}

	userId := user.(token.TokenClaims).UserId
	return userId
}

func handleError(ctx *gin.Context, err error) {
	var alreadyExist *common2.AlreadyExistsError
	var notFound *common2.NotFoundError
	if errors.As(err, &alreadyExist) {
		common2.ConflictError(ctx, err.Error())
	} else if errors.As(err, &notFound) {
		common2.NotFound(ctx, err.Error())
	} else {
		common2.InternalError(ctx, err.Error())
	}
}
