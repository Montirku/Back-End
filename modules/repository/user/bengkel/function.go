package bengkel

import (
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
	"github.com/labstack/echo/v4"
)

func (br *bengkelRepo) GetAllBengkel(bengkels []*bt.Bengkel, offset, pageSize int) ([]*bt.Bengkel, int64, error) {
	var count int64
	bengkel := &bt.Bengkel{}
	if err := br.db.Model(bengkel).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := br.db.
		Preload("BengkelCategory").Preload("BengkelAddress").
		Offset(offset).
		Limit(pageSize).Order("created_at ASC").
		Find(&bengkels).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return bengkels, count, nil
}

func (br *bengkelRepo) GetBengkelById(bengkelId string, bengkel *bt.Bengkel) (*bt.Bengkel, error) {
	if err := br.db.Preload("BengkelCategory").Preload("BengkelAddress").Model(bt.Bengkel{}).Where("id = ?", bengkelId).First(&bengkel).Error; err != nil {
		return bengkel, echo.NewHTTPError(404, err)
	}

	return bengkel, nil
}

func (br *bengkelRepo) GetAllBengkelServices(bengkelId string, bengkelServices []*bt.BengkelServices) ([]*bt.BengkelServices, error) {
	bengkelService := bt.BengkelServices{}
	if err := br.db.Model(bengkelService).Where("bengkel_id = ?", bengkelId).Order("created_at ASC").Find(&bengkelServices).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return bengkelServices, nil
}

func (br *bengkelRepo) GetAllOperationalTime(bengkelId string, operationalTime []*bt.OperationalTime) ([]*bt.OperationalTime, error) {
	if err := br.db.Model(bt.OperationalTime{}).Where("bengkel_id = ?", bengkelId).Order("created_at ASC").Find(&operationalTime).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return operationalTime, nil
}

func (br *bengkelRepo) FilterBengkel(category, status string, offset, pageSize int) ([]*bt.Bengkel, int64, error) {
	var bengkels []*bt.Bengkel
	var count int64

	if err := br.db.Model(&bt.Bengkel{}).
		Where("bengkel_category_id IN (?) AND status LIKE ?",
			br.db.Model(&bt.BengkelCategory{}).Select("id").Where("category LIKE ?", "%"+category+"%"),
			"%"+status+"%").
		Preload("BengkelCategory").Preload("BengkelAddress").
		Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := br.db.Model(&bt.Bengkel{}).
		Where("bengkel_category_id IN (?) AND status LIKE ?",
			br.db.Model(&bt.BengkelCategory{}).Select("id").Where("category LIKE ?", "%"+category+"%"),
			"%"+status+"%").
		Preload("BengkelCategory").Preload("BengkelAddress").
		Order("created_at ASC").
		Offset(offset).Limit(pageSize).Find(&bengkels).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return bengkels, count, nil
}
