package dto

import (
	"errors"
	"time"
)

type Outlet struct {
	ID         *int      `json:"id"`
	Name       string    `json:"name"`
	LocationId *int      `json:"location_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (o *Outlet) Validate() error {
	if o.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

