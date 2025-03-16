package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	common2 "globetrotter/common"
	"globetrotter/game/models"
	"globetrotter/game/service"
	"net/http"
	token "token-manager"
)

type GameHandler struct {
	service service.GameService
}

func NewGameHandler(service service.GameService) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) GetRandomQuestion(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	question, err := h.service.GetRandomQuestion(&requestCtx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, question)
}

func (h *GameHandler) SubmitAnswer(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	var req models.AnswerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common2.BadRequest(ctx, "bad-request", "failed to parse request")
		return
	}
	userId := getUserIdFromContext(ctx)

	response, err := h.service.SubmitAnswer(&requestCtx, userId, &req)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func handleError(ctx *gin.Context, err error) {
	var notFoundErr *common2.NotFoundError
	if errors.As(err, &notFoundErr) {
		common2.NotFound(ctx, err.Error())
	} else {
		common2.InternalError(ctx, err.Error())
	}
}

func getUserIdFromContext(ctx *gin.Context) string {
	user, exists := ctx.Get("user")
	if !exists {
		return ""
	}

	userId := user.(token.TokenClaims).UserId
	return userId
}
