package user

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model   `json:"-"`
	RoleId        uint       `json:"RoleId" form:"RoleId"`
	Email         string     `json:"Email" form:"Email" validate:"required,email"`
	GoogleId      string     `json:"GoogleId" form:"GoogleId"`
	Password      string     `json:"Password" form:"Password" validate:"required,min=8"`
	EmailVerified bool       `json:"EmailVerified"`
	UserDetail    UserDetail `gorm:"foreignKey:UserId"`
}

type RegisterRequest struct {
	FirstName       string `json:"FirstName" form:"FirstName" validate:"required"`
	LastName        string `json:"LastName" form:"LastName" validate:"required"`
	Email           string `json:"Email" form:"Email" validate:"required,email"`
	Password        string `json:"Password" form:"Password" validate:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword" form:"ConfirmPassword" validate:"required,min=8"`
}

type RegisterResponse struct {
	ID        uint   `json:"Id" form:"Id"`
	RoleId    uint   `json:"RoleId" form:"RoleId"`
	FirstName string `json:"FirstName" form:"FirstName" validate:"required"`
	LastName  string `json:"LastName" form:"LastName" validate:"required"`
	Email     string `json:"Email" form:"Email" validate:"required,email"`
	Password  string `json:"Password" form:"Password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required,min=8"`
}

type AuthResponse struct {
	ID           uint   `json:"Id" form:"Id"`
	GoogleId     string `json:"GoogleId" form:"GoogleId"`
	Email        string `json:"Email" form:"Email" validate:"required,email"`
	FirstName    string `json:"FirstName" form:"FirstName"`
	LastName     string `json:"LastName" form:"LastName"`
	Phone        string `json:"Phone" form:"Phone"`
	ProfilePhoto string `json:"ProfilePhoto" form:"ProfilePhoto"`
	AuthToken    string `json:"AuthToken" form:"AuthToken"`
}
