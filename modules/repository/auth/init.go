package auth

import (
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
	"gorm.io/gorm"
)

type AuthRepo interface {
	GetUserByEmail(email string) (*ue.User, error)
	Login(email string) (*ue.AuthResponse, string, uint, bool, error)
	CreateUser(user *ue.RegisterRequest) error
	UserRecovery(userId uint, codeVer string) error
	UpdateUserRecovery(userId uint, codeVer string) error
	GetUserRecovery(userId uint) (ue.UserRecovery, error)
	UpdateEmailVerify(email string) error
	ChangePassword(user ue.RecoveryRequest) error
	DeleteUserRecovery(userId uint) error
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
