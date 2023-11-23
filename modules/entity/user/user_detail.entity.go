package user

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	*gorm.Model  `json:"-"`
	FirstName    string `json:"FirstName" form:"FirstName"`
	LastName     string `json:"LastName" form:"LastName"`
	Phone        string `json:"Phone" form:"Phone" validate:"required,min=10,max=13,numeric"`
	ProfilePhoto string `json:"ProfilePhoto" form:"ProfilePhoto"`
	UserId       uint   `json:"UserId" form:"UserId"`
}
