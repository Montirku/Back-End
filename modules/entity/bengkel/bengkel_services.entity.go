package bengkel

import "gorm.io/gorm"

type BengkelServices struct {
	*gorm.Model
	Name      string  `json:"Name" form:"Name"`
	Price     float64 `json:"Price" form:"Price"`
	BengkelId uint    `json:"BengkelId" form:"BengkelId"`
}

type BengkelServicesResponse struct {
	Name  string
	Price float64
}
