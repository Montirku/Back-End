package seeds

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
)

func CreateBengkelRating() []*bt.BengkelRating {
	BengkelRatings := []*bt.BengkelRating{
		{
			BengkelId: 1,
			UserId:    2,
			Rating:    5.0,
			Comment:   "Bengkelnya mantap, saya suka",
		},
		{
			BengkelId: 1,
			UserId:    4,
			Rating:    3.0,
			Comment:   "Antrinya terlalu lama, saya sampai darah tinggi",
		},
		{
			BengkelId: 1,
			UserId:    5,
			Rating:    5.0,
			Comment:   "Tempat tunggunya nyaman, hasil servis memuaskan",
		},
		{
			BengkelId: 1,
			UserId:    6,
			Rating:    5.0,
			Comment:   "Bengkel langganan, the best bengkel in town",
		},
	}

	return BengkelRatings
}
