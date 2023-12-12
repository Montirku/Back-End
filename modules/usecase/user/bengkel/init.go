package bengkel

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
	br "github.com/fazaalexander/montirku-be/modules/repository/user/bengkel"
)

type BengkelUseCase interface {
	GetAllBengkel(bengkels []*bt.Bengkel, offset, pageSize int) ([]*bt.Bengkel, int64, error)
	GetBengkelById(bengkelId string, bengkel *bt.Bengkel) (*bt.Bengkel, error)
	GetAllBengkelServices(bengkelId string, bengkelServices []*bt.BengkelServices) ([]*bt.BengkelServices, error)
	GetAllOperationalTime(bengkelId string, operationalTime []*bt.OperationalTime) ([]*bt.OperationalTime, error)
	FilterBengkel(category, status string, offset, pageSize int) ([]*bt.Bengkel, int64, error)
}

type bengkelUseCase struct {
	bengkelRepo br.BengkelRepo
}

func New(bengkelRepo br.BengkelRepo) *bengkelUseCase {
	return &bengkelUseCase{
		bengkelRepo,
	}
}
