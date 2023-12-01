package auth

import (
	"errors"

	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func (ar *authRepo) GetUserByEmail(email string) (*ue.User, error) {
	user := &ue.User{}
	err := ar.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("record not found")
	}

	return user, nil
}

func (ar *authRepo) Login(email string) (*ue.AuthResponse, string, uint, bool, error) {
	user := &ue.User{}
	err := ar.db.Preload("UserDetail").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, "", 0, false, errors.New("record not found")
	}

	response := &ue.AuthResponse{
		ID:           user.ID,
		GoogleId:     user.GoogleId,
		Email:        user.Email,
		FirstName:    user.UserDetail.FirstName,
		LastName:     user.UserDetail.LastName,
		Phone:        user.UserDetail.Phone,
		ProfilePhoto: user.UserDetail.ProfilePhoto,
	}

	return response, user.Password, user.RoleId, user.EmailVerified, nil
}

func (ar *authRepo) CreateUser(user *ue.RegisterRequest) error {
	existingUser := ue.User{}
	if err := ar.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already exists")
	}

	userData := ue.User{
		RoleId:        2,
		Email:         user.Email,
		Password:      user.Password,
		EmailVerified: false,
		UserDetail: ue.UserDetail{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
		},
	}

	if err := ar.db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
}
func (ar *authRepo) GetUserRecovery(userId uint) (ue.UserRecovery, error) {
	var recovery ue.UserRecovery
	err := ar.db.Where("user_id = ?", userId).First(&recovery).Error
	if err != nil {
		return recovery, errors.New("record not found")
	}

	return recovery, nil
}

func (ar *authRepo) UserRecovery(userId uint, codeVer string) error {

	userRecover := ue.UserRecovery{
		UserId: userId,
		Code:   codeVer,
	}
	if err := ar.db.Create(&userRecover).Error; err != nil {
		return err
	}

	return nil
}
func (ar *authRepo) UpdateUserRecovery(userId uint, codeVer string) error {

	userRecover := ue.UserRecovery{
		UserId: userId,
		Code:   codeVer,
	}
	result := ar.db.Model(&ue.UserRecovery{}).Where("user_id = ?", userId).Updates(&userRecover)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}

func (ar *authRepo) UpdateEmailVerify(email string) error {

	if err := ar.db.Model(&ue.User{}).Where("email = ?", email).Update("email_verified", true).Error; err != nil {
		return err
	}

	return nil
}

func (ar *authRepo) DeleteUserRecovery(userId uint) error {

	var userRecovery ue.UserRecovery
	result := ar.db.Where("user_id = ?", userId).Delete(&userRecovery)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}

func (ar *authRepo) ChangePassword(user ue.RecoveryRequest) error {
	result := ar.db.Model(&ue.User{}).Where("email = ?", user.Email).Update("password", user.Password)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}
