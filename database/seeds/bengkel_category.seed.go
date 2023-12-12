package seeds

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
)

func CreateBengkelCategory() []*bt.BengkelCategory {
	BengkelCategories := []*bt.BengkelCategory{
		{Category: "Motor"},
		{Category: "Mobil"},
		{Category: "Umum"},
	}

	return BengkelCategories
}
