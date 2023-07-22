package main

import (
	"testing"
	"time"

	"github.com/szymon676/betterdocker/mysql"
)

func TestCreateBooking(t *testing.T) {
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

	// inserting incorrect booking struct
	t.Run("test incorrect booking struct", func(t *testing.T) {
		booking := &Booking{
			DeskID:    12,
			UserID:    91321,
			StartTime: time.Now().Add(time.Second * 10).Format(time.RFC3339),
		}

		err = storage.CreateBooking(booking)
		if err == nil {
			t.Fatal("should return err")
		}
	})

	// inserting correct booking struct
	t.Run("test correct booking struct", func(t *testing.T) {
		booking := &Booking{
			DeskID:    12312,
			UserID:    91321,
			StartTime: time.Now().Add(time.Second * 10).Format(time.RFC3339),
			EndTime:   time.Now().Add(time.Second * 15).Format(time.RFC3339),
		}

		err = storage.CreateBooking(booking)
		if err != nil {
			t.Fatal("shouldn't return err:", err)
		}
	})

	time.Sleep(time.Second * 10)
	container.Stop()
}
