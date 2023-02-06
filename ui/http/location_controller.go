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

type LocationController struct {
	app       application.Application
	responder *infrastructure.Responder
}

func NewLocationController(app application.Application, responder *infrastructure.Responder) *LocationController {
	return &LocationController{app: app, responder: responder}
}

func (ctrl LocationController) GetLocations(ctx *gin.Context) {
	locs, err := ctrl.app.GetRootLocations()
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.SuccessList(len(locs), len(locs), 0, locs))
}

func (ctrl LocationController) CreateLocation(ctx *gin.Context) {
	var createDto dto.LocationDto
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

	location := entity.Location{
		ID:         createDto.ID,
		Name:       createDto.Name,
		Type:       createDto.Type,
		LocationId: createDto.LocationId,
		Children:   nil,
		Outlets:    nil,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
	if len(createDto.Children) > 0 {
		locationsDto, err := ctrl.app.GetLocationsByIds(createDto.Children)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
			return
		}
		location.Children = locationsDto
	}
	_, err = ctrl.app.CreateLocation(&location)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.Success(location))
}

func (ctrl LocationController) GetLocation(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("location_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err.Error()))
		return
	}
	locs, err := ctrl.app.GetLocation(&id)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.Success(locs))
}

func (ctrl LocationController) GetRootLocations(ctx *gin.Context) {
	locations, err := ctrl.app.GetRootLocations()
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, ctrl.responder.SuccessList(len(locations), len(locations), 0, locations))
}
