package repository

import "go-inventory/domain/entity"

type LocationRepository interface {
	GetFirstLevel() ([]*entity.Location, error)
	GetById(id *int) (*entity.Location, error)
	GetByIds(locations []int) ([]*entity.Location, error)
	Save(location *entity.Location) (*entity.Location, error)
}
