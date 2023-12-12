package bengkel

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
)

func (bc *bengkelUseCase) GetAllBengkel(bengkels []*bt.Bengkel, offset, pageSize int) ([]*bt.Bengkel, int64, error) {
	bengkels, count, err := bc.bengkelRepo.GetAllBengkel(bengkels, offset, pageSize)
	return bengkels, count, err
}

func (bc *bengkelUseCase) GetBengkelById(bengkelId string, bengkel *bt.Bengkel) (*bt.Bengkel, error) {
	return bc.bengkelRepo.GetBengkelById(bengkelId, bengkel)
}

func (bc *bengkelUseCase) GetAllBengkelServices(bengkelId string, bengkelServices []*bt.BengkelServices) ([]*bt.BengkelServices, error) {
	return bc.bengkelRepo.GetAllBengkelServices(bengkelId, bengkelServices)
}
func (bc *bengkelUseCase) GetAllOperationalTime(bengkelId string, operationalTime []*bt.OperationalTime) ([]*bt.OperationalTime, error) {
	return bc.bengkelRepo.GetAllOperationalTime(bengkelId, operationalTime)
}
func (bc *bengkelUseCase) FilterBengkel(category, status string, offset, pageSize int) ([]*bt.Bengkel, int64, error) {
	bengkels, count, err := bc.bengkelRepo.FilterBengkel(category, status, offset, pageSize)
	return bengkels, count, err
}
