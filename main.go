package main

import (
	"github.com/szymon676/desk-managment/api"
	"github.com/szymon676/desk-managment/store"
)

const dsn string = "host=localhost port=5432 user=postgres password=1234 dbname=desk-managment sslmode=disable"

func main() {
	db, err := store.NewPostgresDatabase(dsn)
	if err != nil {
		panic(err)
	}
	store := store.NewPostgresStore(db)
	api.Serve(store, "3000")
}
