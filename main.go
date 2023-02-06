package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-inventory/application"
	"go-inventory/domain/entity"
	"go-inventory/infrastructure"
	"go-inventory/infrastructure/repository"
	"go-inventory/infrastructure/service"
	http2 "go-inventory/ui/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("no env gotten")
	}
}

var isDebug = false

func main() {
	db := getDb()
	isDebug = os.Getenv("APP_ENV") != "prod"
	locRepo := repository.NewLocationRepository(db)
	deviceRepo := repository.NewDeviceRepository(db)
	outletRepo := repository.NewOutletRepository(db)
	outletService := service.NewOutletService(locRepo, outletRepo)
	app := application.NewApp(locRepo, deviceRepo, outletService)
	responder := infrastructure.NewResponder()
	locCtrl := http2.NewLocationController(app, responder)
	deviceCtrl := http2.NewDeviceController(app, responder)
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	outletCtrl := http2.NewOutletController(app, responder)
	router := gin.Default()
	router.GET("/locations", locCtrl.GetRootLocations)
	router.POST("/locations", locCtrl.CreateLocation)
	router.GET("/locations/:location_id", locCtrl.GetLocation)
	router.GET("/locations/:location_id/devices", deviceCtrl.GetDevices)
	router.GET("/locations/:location_id/outlets", outletCtrl.GetByLocation)
	router.POST("/devices", deviceCtrl.CreateDevice)
	router.GET("/devices/:device_id", deviceCtrl.GetDeviceById)
	router.POST("/outlets", outletCtrl.CreateOutlet)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, responder.Fail("Not found"))
	})
	log.Fatal(router.Run(":" + os.Getenv("API_PORT")))
}

func getDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	var gormConfig *gorm.Config
	if !isDebug {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		gormConfig = &gorm.Config{}
	}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Panic(err)
	}
	if !isDebug {
		err = Automigrate(db)
		if err != nil {
			log.Panic(err)
		}
	}
	return db
}

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Location{}, &entity.Device{}, &entity.Outlet{})
}
