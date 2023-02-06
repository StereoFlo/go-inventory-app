package http

import (
	"github.com/gin-gonic/gin"
	"go-inventory/application"
	"go-inventory/infrastructure"
	"go-inventory/infrastructure/dto"
	"net/http"
	"strconv"
)

type OutletController struct {
	app       application.Application
	responder *infrastructure.Responder
}

func NewOutletController(
	app application.Application,
	responder *infrastructure.Responder,
) *OutletController {
	return &OutletController{app: app, responder: responder}
}

func (ctrl OutletController) CreateOutlet(ctx *gin.Context) {
	var createDto dto.Outlet
	err := ctx.ShouldBindJSON(&createDto)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err))
		return
	}
	outlet, err := ctrl.app.CreateOutlet(&createDto)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.Success(outlet))
}

func (ctrl OutletController) GetByLocation(ctx *gin.Context) {
	locationId, err := strconv.Atoi(ctx.Param("location_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ctrl.responder.Fail(err))
		return
	}
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	devices, err := ctrl.app.GetOutletsByLocation(locationId)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, ctrl.responder.Fail(err))
		return
	}
	ctx.JSON(http.StatusOK, ctrl.responder.SuccessList(len(devices), limit, offset, devices))
}
