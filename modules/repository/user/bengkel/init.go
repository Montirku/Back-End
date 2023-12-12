package bengkel

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
	"gorm.io/gorm"
)

type BengkelRepo interface {
	GetAllBengkel(bengkels []*bt.Bengkel, offset, pageSize int) ([]*bt.Bengkel, int64, error)
	GetBengkelById(bengkelId string, bengkel *bt.Bengkel) (*bt.Bengkel, error)
	GetAllBengkelServices(bengkelId string, bengkelServices []*bt.BengkelServices) ([]*bt.BengkelServices, error)
	GetAllOperationalTime(bengkelId string, operationalTime []*bt.OperationalTime) ([]*bt.OperationalTime, error)
	FilterBengkel(category, status string, offset, pageSize int) ([]*bt.Bengkel, int64, error)
}

type bengkelRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) BengkelRepo {
	return &bengkelRepo{
		db,
	}
}
