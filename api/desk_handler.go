package api

import (
	"github.com/szymon676/desk-managment/store"
	"github.com/szymon676/desk-managment/types"
	gonimbus "github.com/szymon676/go-nimbus"
)

type ApiServer struct {
	store      store.Store
	listenaddr string
}

func NewApiServer(store store.Store, listenaddr string) *ApiServer {
	return &ApiServer{
		store:      store,
		listenaddr: listenaddr,
	}
}

func (as ApiServer) handleCreateDesk(c gonimbus.Context) error {
	var reqDesk types.BindDesk
	if err := c.BindJSON(&reqDesk); err != nil {
		return err
	}
	if err := as.store.CreateDesk(reqDesk); err != nil {
		return err
	}
	return c.String(200, "desk created successfully")
}

func (as ApiServer) handleGetAvailableDesks(c gonimbus.Context) error {
	desks, err := as.store.GetAvailableDesks()
	if err != nil {
		return err
	}
	if desks == nil {
		return c.String(200, "there are no desks currently available")
	}
	return c.Return(200, "available desks:", desks)
}

func (as ApiServer) handleGetDesks(c gonimbus.Context) error {
	desks, err := as.store.GetAllDesks()
	if err != nil {
		return err
	}
	if desks == nil {
		return c.String(200, "there are no desks available")
	}
	return c.Return(200, "desks:", desks)
}
