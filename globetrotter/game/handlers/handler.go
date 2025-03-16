package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"globetrotter/game/models"
	"globetrotter/game/service"
	"net/http"
	token "token-manager"
	"user-server/common"
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
		common.BadRequest(ctx, "bad-request", "failed to parse request")
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
	var notFoundErr *common.NotFoundError
	if errors.As(err, &notFoundErr) {
		common.NotFound(ctx, err.Error())
	} else {
		common.InternalError(ctx, err.Error())
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
