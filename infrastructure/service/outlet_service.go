package service

import (
	"go-inventory/domain/entity"
	"go-inventory/domain/repository"
	"go-inventory/infrastructure/dto"
	"time"
)

type OutletService struct {
	locationRepo repository.LocationRepository
	outletRepo   repository.OutletRepository
}

func NewOutletService(locationRepo repository.LocationRepository, repo repository.OutletRepository) *OutletService {
	return &OutletService{locationRepo: locationRepo, outletRepo: repo}
}

func (os *OutletService) CreateOutlet(createDto *dto.Outlet) (*entity.Outlet, error) {
	err := createDto.Validate()
	if err != nil {
		return nil, err
	}

	outlet := entity.Outlet{
		Name:      createDto.Name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	if createDto.LocationId != nil {
		_, err = os.locationRepo.GetById(createDto.LocationId)
		if err != nil {
			return nil, err
		}
		outlet.LocationId = createDto.LocationId
	}
	_, err = os.outletRepo.Save(&outlet)
	if err != nil {
		return nil, err
	}

	return &outlet, nil
}

func (os *OutletService) GetByLocation(locationId int) ([]*entity.Outlet, error) {
	return os.outletRepo.GetByLocationId(locationId)
}
