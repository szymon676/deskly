package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/szymon676/betterdocker/mysql"
)

func TestHandleCreateBooking(t *testing.T) {
	opts := &mysql.MySQLContainerOptions{
		Database:     "testing",
		RootPassword: "1234",
	}

	container := mysql.NewMySQLContainer(opts)
	err := container.Run()
	if err != nil {
		t.Fatal("failed to init testing container")
	}

	dsn := "root:1234@tcp(127.0.0.1:3306)/testing"
	storage, err := NewStorage(dsn)
	if err != nil {
		panic(err)
	}

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
	resp, _ := app.Test(req, 1)
	if resp.StatusCode != http.StatusOK {
		t.Fatal("expected code 200 not", resp.StatusCode)
	}

	container.Stop()
}
