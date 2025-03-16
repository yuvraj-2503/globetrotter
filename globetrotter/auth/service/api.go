package service

import "context"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthService interface {
	Login(ctx *context.Context, request *LoginRequest) (string, error)
	SignUp(ctx *context.Context, request *SignUpRequest) (string, error)
}
