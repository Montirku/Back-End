package bengkel

import "gorm.io/gorm"

type BengkelAddress struct {
	*gorm.Model  `json:"-"`
	DistrictName string  `json:"DistrictName" form:"DistrictName"`
	CityName     string  `json:"CityName" form:"CityName"`
	Address      string  `json:"Address" form:"Address"`
	Latitude     float64 `json:"Latitude" form:"Latitude"`
	Longitude    float64 `json:"Longitude" form:"Longitude"`
	BengkelId    uint    `json:"BengkelId" form:"BengkelId"`
}

type BengkelAddressResponse struct {
	DistrictName string
	CityName     string
	Address      string
	Latitude     float64
	Longitude    float64
}
