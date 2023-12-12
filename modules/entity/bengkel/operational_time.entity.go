package bengkel

import (
	"gorm.io/gorm"
)

type OperationalTime struct {
	*gorm.Model
	Day       string `json:"Day" form:"Day"`
	OpenTime  string `json:"OpenTime" form:"OpenTime"`
	CloseTime string `json:"CloseTime" form:"CloseTime"`
	BengkelId uint   `json:"BengkelId" form:"BengkelId"`
}

type OperationalTimeResponse struct {
	Day       string
	OpenTime  string
	CloseTime string
}
