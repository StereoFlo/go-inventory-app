package entity

import (
	"time"
)

type Device struct {
	ID          *int      `gorm:"primaryKey;auto_increment" json:"id"`
	Name        string    `json:"name"`
	NetName     *string   `json:"net_name" gorm:"index:idx_member"`
	IP          *string   `json:"ip" gorm:"index:idx_member"`
	TimeToCheck int       `json:"time_to_check"`
	LocationId  *int      `json:"location_id" gorm:"index"`
	Outlets     []*Outlet `gorm:"many2many:outlets_devices;" json:"outlets"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
