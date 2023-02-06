package service

import (
	"go-inventory/domain/entity"
	"go-inventory/infrastructure/dto"
)

type OutletServiceInterface interface {
	CreateOutlet(createDto *dto.Outlet) (*entity.Outlet, error)
	GetByLocation(locationId int) ([]*entity.Outlet, error)
}
