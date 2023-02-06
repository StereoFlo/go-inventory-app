package http

import (
	"github.com/gin-gonic/gin"
	"go-inventory/application"
	"go-inventory/domain/entity"
	"go-inventory/infrastructure"
	"go-inventory/infrastructure/dto"
	"net/http"
	"strconv"
	"time"
)

type DeviceController struct {
	app       application.Application
	responder *infrastructure.Responder
}

func NewDeviceController(app application.Application, responder *infrastructure.Responder) *DeviceController {
	return &DeviceController{app: app, responder: responder}
}

func (ctrl DeviceController) GetDevices(ctx *gin.Context) {
	locationId, err := strconv.Atoi(ctx.Param("location_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}
	if limit == 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}
	devices, err := ctrl.app.GetDeviceList(locationId, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.SuccessList(len(devices), limit, offset, devices))
}

func (ctrl DeviceController) GetDeviceById(ctx *gin.Context) {
	deviceId, err := strconv.Atoi(ctx.Param("device_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}
	device, err := ctrl.app.GetDeviceById(deviceId)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.Success(device))
}

func (ctrl DeviceController) CreateDevice(ctx *gin.Context) {
	var createDto dto.DeviceDto
	err := ctx.ShouldBindJSON(&createDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}
	err = createDto.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}

	deviceByName, err := ctrl.app.GetDeviceByNameAndIp(*createDto.NetName, *createDto.IP)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}

	if deviceByName.IP != nil &&
		createDto.IP != nil &&
		createDto.NetName != nil &&
		deviceByName.NetName != nil &&
		*deviceByName.IP == *createDto.IP &&
		*deviceByName.NetName == *createDto.NetName {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail("device exists"))
		return
	}

	device := entity.Device{
		Name:      createDto.Name,
		NetName:   createDto.NetName,
		IP:        createDto.IP,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	if createDto.LocationId != nil {
		lid := *createDto.LocationId
		locationsDto, err := ctrl.app.GetLocation(&lid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
			return
		}
		device.LocationId = locationsDto.ID
	}
	_, err = ctrl.app.CreateDevice(&device)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.Success(device))
}
