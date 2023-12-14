package bengkel

import (
	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	"gorm.io/gorm"
)

type BengkelServices struct {
	*gorm.Model
	Name               string                 `json:"Name" form:"Name"`
	Price              float64                `json:"Price" form:"Price"`
	BengkelId          uint                   `json:"BengkelId" form:"BengkelId"`
	TransactionDetails []te.TransactionDetail `gorm:"foreignKey:ServiceId"`
}

type BengkelServicesResponse struct {
	ServiceId string
	Name  string
	Price float64
}
