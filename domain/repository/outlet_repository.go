package repository

import "go-inventory/domain/entity"

type OutletRepository interface {
	GetById(id int) (*entity.Outlet, error)
	GetByIds(outlet []int) ([]*entity.Outlet, error)
	GetByLocationId(locationId int) ([]*entity.Outlet, error)
	Save(outlet *entity.Outlet) (*entity.Outlet, error)
}
