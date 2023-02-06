package repository

import (
	"go-inventory/domain/entity"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (dr DeviceRepository) GetByLocationId(locationId int, limit int, offset int) ([]*entity.Device, error) {
	var devices []*entity.Device
	err := dr.db.
		Debug().
		Preload("Outlets").
		Where("location_id", locationId).
		Limit(limit).
		Offset(offset).
		Find(&devices).Error

	return devices, err
}

func (dr DeviceRepository) GetById(id int) (*entity.Device, error) {
	var device1 *entity.Device
	err := dr.db.Debug().Where("id = ?", id).First(&device1).Error

	return device1, err
}

func (dr DeviceRepository) GetByNameAndIp(name string, ip string) (*entity.Device, error) {
	var device1 *entity.Device
	err := dr.db.Debug().Where("net_name like ? and ip = ?", "%"+name+"%", ip).Find(&device1).Error

	return device1, err
}

func (dr DeviceRepository) Save(device *entity.Device) (*entity.Device, error) {
	err := dr.db.Debug().Create(&device).Error
	if err != nil {
		return nil, err
	}

	return device, nil
}
