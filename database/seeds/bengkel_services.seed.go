package seeds

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
)

func CreateBengkelServices() []*bt.BengkelServices {
	BengkelServices := []*bt.BengkelServices{
		{
			BengkelId: 1,
			Name:      "Service Umum",
			Price:     150000,
		},
		{
			BengkelId: 1,
			Name:      "Service Berkala",
			Price:     150000,
		},
		{
			BengkelId: 1,
			Name:      "Ganti Oli",
			Price:     60000,
		},
		{
			BengkelId: 1,
			Name:      "Ganti Aki",
			Price:     50000,
		},
		{
			BengkelId: 1,
			Name:      "Ganti Kampas Rem",
			Price:     30000,
		},
		{
			BengkelId: 1,
			Name:      "Ganti Lampu",
			Price:     50000,
		},
		{
			BengkelId: 1,
			Name:      "Ganti Ban",
			Price:     150000,
		},
	}

	return BengkelServices
}
