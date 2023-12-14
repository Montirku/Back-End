package transaction

import (
	"gorm.io/gorm"
)

type TransactionDetail struct {
	*gorm.Model
	TransactionId uint
	BengkelName   string  `json:"BengkelName" form:"BengkelName" validate:"required"`
	ServiceId     string  `json:"ServiceId" form:"ServiceId"`
	ServiceName   string  `json:"ServiceName" form:"ServiceName" validate:"required"`
	Qty           uint    `json:"Qty" form:"Qty" validate:"required"`
	SubTotalPrice float64 `json:"SubTotalPrice" form:"SubTotalPrice" validate:"required"`
}
