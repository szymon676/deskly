package store

import (
	"github.com/szymon676/desk-managment/types"
)

type Store interface {
	GetAvailableDesks() ([]types.Desk, error)
	CreateDesk(desk types.BindDesk) error
}
