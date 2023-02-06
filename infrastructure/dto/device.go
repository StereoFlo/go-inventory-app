package dto

import (
	"errors"
)

type DeviceDto struct {
	ID          *int    `json:"id"`
	Name        string  `json:"name"`
	NetName     *string `json:"net_name" gorm:"index"`
	IP          *string `json:"ip"`
	TimeToCheck int     `json:"time_to_check"`
	LocationId  *int    `json:"location_id"`
}

func (d *DeviceDto) Validate() error {
	if d.Name == "" {
		return errors.New("name cannot be empty")
	}
	if d.NetName == nil {
		return errors.New("net_name cannot be empty")
	}
	if d.IP == nil {
		return errors.New("ip cannot be empty")
	}
	return nil
}
