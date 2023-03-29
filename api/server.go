package api

import (
	"github.com/szymon676/desk-managment/store"
	gonimbus "github.com/szymon676/go-nimbus"
)

func Serve(store store.Store, listenaddr string) {
	g := gonimbus.New()
	as := NewApiServer(store, listenaddr)

	g.Post("/desks", makeHTTPHandler(as.handleCreateDesk))
	g.Get("/available/desks", makeHTTPHandler(as.handleGetAvailableDesks))
	g.Get("/desks", makeHTTPHandler(as.handleGetDesks))

	g.Serve(as.listenaddr)
}

type apiFunc func(c gonimbus.Context) error

func makeHTTPHandler(fn apiFunc) gonimbus.HandlerFunc {
	return func(c gonimbus.Context) {
		if err := fn(c); err != nil {
			c.Return(404, "error : ", err)
		}
	}
}
