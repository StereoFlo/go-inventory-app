package dto

import (
	"errors"
)

type LocationDto struct {
	ID         *int   `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	LocationId *int   `json:"location_id"`
	Children   []int  `json:"children"`
	Outlets    []int  `json:"outlets"`
}

func (loc *LocationDto) Validate() error {
	if loc.Name == "" {
		return errors.New("name is empty")
	}
	if loc.Type == "" {
		return errors.New("type is empty")
	}

	return nil
}
