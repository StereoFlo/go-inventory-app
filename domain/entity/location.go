package entity

import (
	"time"
)

type Location struct {
	ID         *int        `gorm:"primaryKey;auto_increment" json:"id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	LocationId *int        `json:"location_id"`
	Children   []*Location `json:"children"`
	Devices    []*Device   `json:"devices"`
	Outlets    []*Outlet   `json:"outlets"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
