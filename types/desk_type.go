package types

import (
	"errors"
	"time"
)

type Desk struct {
	ID           int    `json:"id"`
	Location     string `json:"location"`
	Availability bool   `json:"availability"`
}

type Booking struct {
	ID       int       `json:"id"`
	Location int       `json:"location"`
	UserID   int       `json:"userID"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

type BindDesk struct {
	Location uint16 `json:"location"`
}

func NewDeskFromRequest(bd BindDesk) error {
	if bd.Location < 1 {
		return errors.New("location must be greater than 1 :D")
	}
	return nil
}
