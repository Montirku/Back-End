package transaction

import (
	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model

	UserId             uint `validate:"required"`
	StatusTransaction  string
	ReceiptNumber      string
	TransactionId      string `validate:"required"`
	PaymentMethod      string
	PaymentStatus      string
	PaymentUrl         string `validate:"required"`
	CanceledReason     string
	TotalPrice         float64             `validate:"required"`
	TransactionDetails []TransactionDetail `json:"TransactionDetails" form:"TransactionDetails" gorm:"foreignKey:TransactionId"`
}
