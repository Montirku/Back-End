package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	ev "github.com/fazaalexander/montirku-be/helper/emailverification"
	pw "github.com/fazaalexander/montirku-be/helper/password"
	vld "github.com/fazaalexander/montirku-be/helper/validator"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func (ac *authUsecase) Register(request *ue.RegisterRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
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

func (ac *authUsecase) EmailVerification(request *ue.VerifyEmailRequest) (string, error) {

	if err := vld.Validation(request); err != nil {
		return "", err
	}

	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return "", errors.New("email tidak ditemukan")
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
		return errors.New("email tidak ditemukan")
	}
	userRecovery, err := ac.authRepo.GetUserRecovery(user.ID)
	if err != nil {
		return errors.New("kode verifikasi tidak ditemukan")
	}

	expTime := userRecovery.CreatedAt.Add(15 * time.Minute)

	if !time.Now().Before(expTime) {
		return errors.New("kode otp kadaluarsa")
	}

	if request.CodeOtp != userRecovery.Code {
		return errors.New("kode verifikasi salah")
	}

	err = ac.authRepo.UpdateEmailVerify(request.Email)
	if err != nil {
		return err
	}

	return nil
}
