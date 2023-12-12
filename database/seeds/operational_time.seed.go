package seeds

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
)

func CreateOperationalTime() []*bt.OperationalTime {
	openTime := "07:00"
	closeTime := "17:00"
	OperationalTimes := []*bt.OperationalTime{
		{
			BengkelId: 1,
			Day:       "Senin",
			OpenTime:  openTime,
			CloseTime: closeTime,
		},
		{
			BengkelId: 1,
			Day:       "Selasa",
			OpenTime:  openTime,
			CloseTime: closeTime,
		},
		{
			BengkelId: 1,
			Day:       "Rabu",
			OpenTime:  openTime,
			CloseTime: closeTime,
		},
		{
			BengkelId: 1,
			Day:       "Kamis",
			OpenTime:  openTime,
			CloseTime: closeTime,
		},
		{
			BengkelId: 1,
			Day:       "Jumat",
			OpenTime:  openTime,
			CloseTime: closeTime,
		},
		{
			BengkelId: 1,
			Day:       "Sabtu",
			OpenTime:  openTime,
			CloseTime: closeTime,
		},
		{
			BengkelId: 1,
			Day:       "Minggu",
			OpenTime:  openTime,
			CloseTime: closeTime,
		},
	}

	return OperationalTimes
}
