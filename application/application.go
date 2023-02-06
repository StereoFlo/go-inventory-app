package application

import (
	"go-inventory/domain/entity"
	"go-inventory/domain/repository"
	"go-inventory/domain/service"
	"go-inventory/infrastructure/dto"
)

type Application interface {
	GetLocation(id *int) (*entity.Location, error)
	CreateLocation(loc *entity.Location) (*entity.Location, error)
	GetRootLocations() ([]*entity.Location, error)
	GetDeviceById(id int) (*entity.Device, error)
	GetDeviceByNameAndIp(name string, ip string) (*entity.Device, error)
	GetDeviceList(locationId int, limit int, offset int) ([]*entity.Device, error)
	GetLocationsByIds(children []int) ([]*entity.Location, error)
	CreateDevice(device *entity.Device) (*entity.Device, error)
	CreateOutlet(outlet *dto.Outlet) (*entity.Outlet, error)
	GetOutletsByLocation(locationId int) ([]*entity.Outlet, error)
}

type App struct {
	locationRepo repository.LocationRepository
	deviceRepo   repository.DeviceRepository
	outletRepo   service.OutletServiceInterface
}

func NewApp(
	locationRepo repository.LocationRepository,
	deviceRepo repository.DeviceRepository,
	outletRepo service.OutletServiceInterface,
) Application {
	return &App{locationRepo: locationRepo, deviceRepo: deviceRepo, outletRepo: outletRepo}
}

func (app App) CreateLocation(loc *entity.Location) (*entity.Location, error) {
	return app.locationRepo.Save(loc)
}

func (app App) GetLocation(id *int) (*entity.Location, error) {
	return app.locationRepo.GetById(id)
}

func (app App) GetRootLocations() ([]*entity.Location, error) {
	return app.locationRepo.GetFirstLevel()
}

func (app App) GetLocationsByIds(children []int) ([]*entity.Location, error) {
	return app.locationRepo.GetByIds(children)
}

func (app App) GetDeviceById(id int) (*entity.Device, error) {
	return app.deviceRepo.GetById(id)
}

func (app App) GetDeviceByNameAndIp(name string, ip string) (*entity.Device, error) {
	return app.deviceRepo.GetByNameAndIp(name, ip)
}

func (app App) GetDeviceList(locationId int, limit int, offset int) ([]*entity.Device, error) {
	return app.deviceRepo.GetByLocationId(locationId, limit, offset)
}

func (app App) CreateDevice(device *entity.Device) (*entity.Device, error) {
	return app.deviceRepo.Save(device)
}

func (app App) CreateOutlet(outlet *dto.Outlet) (*entity.Outlet, error) {
	return app.outletRepo.CreateOutlet(outlet)
}

func (app App) GetOutletsByLocation(locationId int) ([]*entity.Outlet, error) {
	return app.outletRepo.GetByLocation(locationId)
}
