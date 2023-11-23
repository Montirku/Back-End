package auth

import (
	ac "github.com/fazaalexander/montirku-be/modules/usecase/auth"
)

type AuthHandler struct {
	authUsecase ac.AuthUsecase
}

func New(authUsecase ac.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase,
	}
}
