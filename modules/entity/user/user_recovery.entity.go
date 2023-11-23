package user

import "time"

type UserRecovery struct {
	ID        uint `gorm:"primarykey"`
	UserId    uint
	Code      string
	CreatedAt time.Time
}

type VerifyEmailRequest struct {
	Email string `json:"Email" form:"Email" validate:"required,email"`
}

type ForgotPasswordRequest struct {
	Email string `json:"Email" form:"Email" validate:"required,email"`
}

type VerifOtp struct {
	Email   string `json:"Email" form:"Email" validate:"required,email"`
	CodeOtp string `json:"CodeOtp" form:"CodeOtp" validate:"required"`
}

type RecoveryRequest struct {
	Email           string `json:"Email" form:"Email" validate:"required,email"`
	Password        string `json:"Password" form:"Password" validate:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword" form:"Password" validate:"required,min=8"`
}
