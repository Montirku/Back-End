package auth

import (
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
	ar "github.com/fazaalexander/montirku-be/modules/repository/auth"
)

type AuthUsecase interface {
	Register(user *ue.RegisterRequest) error
	EmailVerification(request *ue.VerifyEmailRequest) (string, error)
	EmailOTPVerification(request ue.VerifOtp) error
}

type authUsecase struct {
	authRepo ar.AuthRepo
}

func New(adminRepo ar.AuthRepo) *authUsecase {
	return &authUsecase{
		adminRepo,
	}
}
