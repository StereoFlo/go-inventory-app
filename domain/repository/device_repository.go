package repository

import "go-inventory/domain/entity"

type DeviceRepository interface {
	GetByLocationId(locationId int, limit int, offset int) ([]*entity.Device, error)
	GetById(id int) (*entity.Device, error)
	GetByNameAndIp(name string, ip string) (*entity.Device, error)
	Save(device *entity.Device) (*entity.Device, error)
}
