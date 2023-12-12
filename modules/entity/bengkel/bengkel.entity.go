package bengkel

import (
	"gorm.io/gorm"
)

type Bengkel struct {
	*gorm.Model       `json:"-"`
	Name              string  `json:"Name" form:"Name"`
	PhoneNumber       string  `json:"PhoneNumber" form:"PhoneNumber"`
	StartingPrice     float64 `json:"StartingPrice" form:"StartingPrice"`
	BengkelPhotoUrl   string  `json:"BengkelPhotoUrl" form:"BengkelPhotoUrl"`
	AvgRating         float64 `json:"AvgRating" form:"AvgRating"`
	Status            string  `json:"Status" form:"Status"`
	UserId            uint    `json:"UserId" form:"UserId"`
	BengkelCategoryId uint    `json:"BengkelCategoryId" form:"BengkelCategoryId"`
	BengkelCategory   BengkelCategory
	BengkelAddress    BengkelAddress    `gorm:"foreignKey:BengkelId"`
	BengkelServices   []BengkelServices `gorm:"foreignKey:BengkelId"`
	OperationalTime   []OperationalTime `gorm:"foreignKey:BengkelId"`
	BengkelRating     []BengkelRating   `gorm:"foreignKey:BengkelId"`
}

type BengkelResponse struct {
	Category        string
	Name            string
	PhoneNumber     string
	StartingPrice   float64
	BengkelPhotoUrl string
	AvgRating       float64
	Status          string
	DistrictName    string
	CityName        string
}

type BengkelDetailResponse struct {
	Category        string
	Name            string
	PhoneNumber     string
	StartingPrice   float64
	BengkelPhotoUrl string
	AvgRating       float64
	Status          string
	Address         BengkelAddressResponse
	Services        []BengkelServicesResponse
	OperationalTime []OperationalTimeResponse
}
