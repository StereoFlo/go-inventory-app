package repository

import (
	"go-inventory/domain/entity"
	"gorm.io/gorm"
)

type OutletRepository struct {
	db *gorm.DB
}

func NewOutletRepository(db *gorm.DB) *OutletRepository {
	return &OutletRepository{db: db}
}

func (ol *OutletRepository) GetById(id int) (*entity.Outlet, error) {
	var o *entity.Outlet
	err := ol.db.Where("id = ?", id).First(&o).Error

	return o, err
}

func (ol *OutletRepository) GetByIds(outletIds []int) ([]*entity.Outlet, error) {
	var o []*entity.Outlet
	err := ol.db.Debug().Where(outletIds).Find(&o, outletIds).Error
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (ol *OutletRepository) GetByLocationId(locationId int) ([]*entity.Outlet, error) {
	var o []*entity.Outlet
	err := ol.db.
		Debug().
		Preload("Devices").
		Where("location_id = ?", locationId).
		Find(&o).
		Error
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (ol *OutletRepository) Save(outlet *entity.Outlet) (*entity.Outlet, error) {
	err := ol.db.Debug().Create(&outlet).Error
	if err != nil {
		return nil, err
	}

	return outlet, nil
}
