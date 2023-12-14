package order

import (
	"time"
)

type CanceledOrder struct {
	TransactionId  string `json:"TransactionId" form:"TransactionId" validate:"required"`
	CanceledReason string `json:"CanceledReason" form:"CanceledReason" validate:"required"`
}
type ConfirmOrder struct {
	TransactionId string `json:"TransactionId" form:"TransactionId" validate:"required"`
}

type Order struct {
	TransactionId      string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	AddressId          uint
	StatusTransaction  string
	ReceiptNumber      string
	TotalProductPrice  float64
	TotalShippingPrice float64
	Point              float64
	PaymentMethod      string
	PaymentStatus      string
	ExpeditionName     string
	VoucherId          uint
	Discount           float64
	TotalPrice         float64
	EstimationDay      string
	PaymentUrl         string
	ExpeditionRating   float32
	CanceledReason     string
	OrderDetail        []OrderDetail
}

type OrderDetail struct {
	ServiceId       string
	Qty             uint
	SubTotalPrice   float64
	ServiceName     string
	ProductImageUrl string
	RatingProductId uint
}
