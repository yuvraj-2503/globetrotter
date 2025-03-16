package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"globetrotter/auth/service"
	common2 "globetrotter/common"
	"net/http"
)

type AuthenticationHandler struct {
	authService service.AuthService
}

func NewAuthenticationHandler(authService service.AuthService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authService: authService,
	}
}

func (a *AuthenticationHandler) Login(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	var loginRequest service.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		common2.BadRequest(ctx, "bad-request", "Invalid request")
		return
	}
	result, err := a.authService.Login(&requestCtx, &loginRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": result,
	})
}

func (a *AuthenticationHandler) SignUp(ctx *gin.Context) {
	requestCtx := ctx.Request.Context()
	var signupRequest service.SignUpRequest
	if err := ctx.ShouldBindJSON(&signupRequest); err != nil {
		common2.BadRequest(ctx, "bad-request", "Invalid request")
		return
	}
	result, err := a.authService.SignUp(&requestCtx, &signupRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": result,
	})
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
