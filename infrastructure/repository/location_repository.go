package repository

import (
	"go-inventory/domain/entity"
	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (lr LocationRepository) GetFirstLevel() ([]*entity.Location, error) {
	var locations []*entity.Location
	err := lr.db.Debug().Where("location_id is null").Find(&locations).Error

	return locations, err
}

func (lr LocationRepository) GetById(id *int) (*entity.Location, error) {
	var location1 *entity.Location
	err := lr.db.
		Debug().
		Preload("Children").
		Preload("Devices").
		Preload("Outlets").
		Where("id = ?", id).
		First(&location1).
		Error

	return location1, err
}

func (lr LocationRepository) Save(location *entity.Location) (*entity.Location, error) {
	err := lr.db.Debug().Create(&location).Error
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (lr LocationRepository) GetByIds(locationsIds []int) ([]*entity.Location, error) {
	var locations []*entity.Location
	err := lr.db.Debug().Where(locationsIds).Find(&locations, locationsIds).Error
	if err != nil {
		return nil, err
	}

	return locations, nil
}
