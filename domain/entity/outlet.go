package entity

import (
	"time"
)

type Outlet struct {
	ID         *int      `gorm:"primaryKey;auto_increment" json:"id"`
	Name       string    `json:"name"`
	LocationId *int      `json:"location_id"`
	Devices    []*Outlet `gorm:"many2many:outlets_devices;" json:"devices"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
