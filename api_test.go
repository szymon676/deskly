package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/szymon676/betterdocker/mysql"
)

func TestHandleCreateBooking(t *testing.T) {
	// Set up the database container
	opts := &mysql.MySQLContainerOptions{
		Database:     "testing",
		RootPassword: "1234",
	}

	container := mysql.NewMySQLContainer(opts)
	err := container.Run()
	if err != nil {
		t.Fatal("failed to init testing container")
	}

	defer container.Stop()

	dsn := "root:1234@tcp(127.0.0.1:3306)/testing"
	storage, err := NewStorage(dsn)
	if err != nil {
		panic(err)
	}

	tx := storage.db.Begin()
	defer tx.Rollback()

	as := apiserver{storage: storage}
	app := as.SetupServer()

	bookingReq := &BookingRequest{
		DeskID:    123,
		UserID:    321,
		StartTime: time.Now().Add(time.Second * 15).Format(time.ANSIC),
		EndTime:   time.Now().Add(time.Second * 30).Format(time.ANSIC),
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(bookingReq)
	if err != nil {
		t.Fatal("failed to encode body:", err)
	}

	req := httptest.NewRequest("POST", "/booking", &b)
	resp, err := app.Test(req, 1)
	if err != nil {
		t.Fatal("failed to send request:", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected code 200")
}
