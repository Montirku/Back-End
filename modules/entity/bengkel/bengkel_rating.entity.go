package bengkel

import "gorm.io/gorm"

type BengkelRating struct {
	*gorm.Model `json:"-"`
	Rating      float64 `json:"Rating" form:"Rating"`
	Comment     string  `json:"Comment" form:"Comment"`
	UserId      uint    `json:"UserId" form:"UserId"`
	BengkelId   uint    `json:"BengkelId" form:"BengkelId"`
	// TransactionDetailId uint    `json:"TransactionDetailId" form:"TransactionDetailId"`
}

// type GetAllBengkelRatingResponse struct {
// 	BengkelId uint
// 	UserId    uint
// 	UserName string
// 	UserProfilePhoto string
// 	BengkelName string
// 	BengkelCategory string
// 	Rating    float64
// 	Comment   string
// }
