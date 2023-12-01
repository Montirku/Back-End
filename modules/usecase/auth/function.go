package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	ev "github.com/fazaalexander/montirku-be/helper/emailverification"
	fp "github.com/fazaalexander/montirku-be/helper/forgotpassword"
	pw "github.com/fazaalexander/montirku-be/helper/password"
	vld "github.com/fazaalexander/montirku-be/helper/validator"
	"github.com/fazaalexander/montirku-be/middleware/jwt"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func (ac *authUsecase) Register(request *ue.RegisterRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	if request.Password != request.ConfirmPassword {
		return errors.New("password doesn't match")
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = string(hashedPassword)

	err = ac.authRepo.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) Login(request *ue.LoginRequest) (interface{}, uint, error) {
	if err := vld.Validation(request); err != nil {
		return nil, 0, err
	}

	response, password, role, emailVerified, err := ac.authRepo.Login(request.Email)

	if err != nil {
		return nil, 0, errors.New("invalid email or password")
	}

	err = pw.VerifyPassword(password, request.Password)
	if err != nil {
		return nil, 0, errors.New("invalid email or password")
	}

	if !emailVerified {
		return nil, 0, errors.New("email not verified")
	}

	token, err := jwt.CreateToken(int(response.ID), response.Email)
	if err != nil {
		return nil, 0, err
	}
	response.AuthToken = token

	return response, role, nil
}

func (ac *authUsecase) EmailVerification(request *ue.VerifyEmailRequest) (string, error) {

	if err := vld.Validation(request); err != nil {
		return "", err
	}

	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	codeVerRandRune := make([]rune, 6)
	for i := 0; i < 6; i++ {
		codeVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes))]
	}
	codeVerPassword := string(codeVerRandRune)
	fmt.Println("Generated Code:", codeVerPassword)

	_, err = ac.authRepo.GetUserRecovery(user.ID)
	if err != nil {
		err = ac.authRepo.UserRecovery(user.ID, codeVerPassword)
		if err != nil {
			return "", err
		}
	} else {
		err = ac.authRepo.UpdateUserRecovery(user.ID, codeVerPassword)
		if err != nil {
			return "", err
		}
	}

	err = ev.EmailVerification(request.Email, codeVerPassword)
	if err != nil {
		return "", err
	}

	return request.Email, nil
}
func (ac *authUsecase) EmailOTPVerification(request ue.VerifOtp) error {
	if err := vld.Validation(request); err != nil {
		return err
	}
	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return errors.New("email not found")
	}
	userRecovery, err := ac.authRepo.GetUserRecovery(user.ID)
	if err != nil {
		return errors.New("verification code not found")
	}

	expTime := userRecovery.CreatedAt.Add(15 * time.Minute)

	if !time.Now().Before(expTime) {
		return errors.New("otp code expired")
	}

	if request.CodeOtp != userRecovery.Code {
		return errors.New("incorrect otp verification code")
	}

	err = ac.authRepo.UpdateEmailVerify(request.Email)
	if err != nil {
		return err
	}

	err = ac.authRepo.DeleteUserRecovery(user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) ForgotPassword(request *ue.ForgotPasswordRequest) (string, error) {
	if err := vld.Validation(request); err != nil {
		return "", err
	}

	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	codeVerRandRune := make([]rune, 6)
	for i := 0; i < 6; i++ {
		codeVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes))]
	}
	codeVerPassword := string(codeVerRandRune)
	fmt.Println("Generated Code:", codeVerPassword)

	_, err = ac.authRepo.GetUserRecovery(user.ID)
	if err != nil {
		err = ac.authRepo.UserRecovery(user.ID, codeVerPassword)
		if err != nil {
			return "", err
		}
	} else {
		err = ac.authRepo.UpdateUserRecovery(user.ID, codeVerPassword)
		if err != nil {
			return "", err
		}
	}

	err = fp.ForgotPassword(request.Email, codeVerPassword)
	if err != nil {
		return "", err
	}

	return request.Email, nil
}
func (ac *authUsecase) PasswordOTPVerification(request ue.VerifOtp) error {
	if err := vld.Validation(request); err != nil {
		return err
	}
	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return errors.New("email not found")
	}
	userRecovery, err := ac.authRepo.GetUserRecovery(user.ID)
	if err != nil {
		return errors.New("verification code not found")
	}

	expTime := userRecovery.CreatedAt.Add(15 * time.Minute)

	if !time.Now().Before(expTime) {
		return errors.New("otp code expired")
	}

	if request.CodeOtp != userRecovery.Code {
		return errors.New("incorrect otp verification code")
	}

	return nil
}
func (ac *authUsecase) ChangePassword(request ue.RecoveryRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return errors.New("email not found")
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}
	request.Password = string(hashedPassword)
	err = ac.authRepo.ChangePassword(request)
	if err != nil {
		return err
	}
	err = ac.authRepo.DeleteUserRecovery(user.ID)
	if err != nil {
		return err
	}
	return nil
}
