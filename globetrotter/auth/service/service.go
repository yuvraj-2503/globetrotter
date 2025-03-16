package service

import (
	"context"
	"globetrotter/auth/db"
	"globetrotter/auth/models"
	"globetrotter/common"
	"time"
	token "token-manager"
)

type AuthServiceImpl struct {
	userDB       db.UserDB
	tokenManager token.TokenManager
}

func NewAuthServiceImpl(userDB db.UserDB, tokenManager token.TokenManager) *AuthServiceImpl {
	return &AuthServiceImpl{
		userDB:       userDB,
		tokenManager: tokenManager,
	}
}

func (a *AuthServiceImpl) Login(ctx *context.Context, request *LoginRequest) (string, error) {
	user, err := a.userDB.GetByEmail(ctx, request.Email)
	if err != nil {
		return "", err
	}

	if user.Password != request.Password {
		return "", &common.InvalidPasswordError{Message: "Password does not match"}
	}

	claims := getClaims(request.Email)
	jwt, err := a.tokenManager.Generate(claims)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (a *AuthServiceImpl) SignUp(ctx *context.Context, request *SignUpRequest) (string, error) {
	err := a.userDB.InsertUser(ctx, &models.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return "", err
	}

	claims := getClaims(request.Email)
	jwt, err := a.tokenManager.Generate(claims)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func getClaims(email string) *token.TokenClaims {
	iat := time.Now()
	exp := iat.Add(time.Hour * 24 * 7)
	var claims = &token.TokenClaims{
		UserId:    email,
		EmailId:   email,
		MachineId: "",
		App:       "globetrotter",
		IAT:       &iat,
		EXP:       &exp,
		Kind:      "USER",
		Sub:       email,
	}
	return claims
}
