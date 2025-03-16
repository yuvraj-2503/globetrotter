package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	common2 "globetrotter/common"
	"globetrotter/user/db"
	"globetrotter/user/service"
	"net/http"
	token "token-manager"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) RegisterUser(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	var userDetails db.UserDetails
	username := ctx.Query("username")

	userId := getUserIdFromContext(ctx)
	userDetails.UserId = userId
	userDetails.Username = username
	err := u.userService.RegisterUser(&requestCtx, &userDetails)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (u *UserHandler) GetUserByUsername(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	username := ctx.Query("username")

	result, err := u.userService.GetUserByUsername(&requestCtx, username)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response := &UserResponse{
		Username: result.Username,
		Score:    result.Score,
	}
	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) GetUserById(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	userId := getUserIdFromContext(ctx)

	result, err := u.userService.GetUserById(&requestCtx, userId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response := &UserResponse{
		Username: result.Username,
		Score:    result.Score,
	}
	ctx.JSON(http.StatusOK, response)
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

type UserResponse struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}
