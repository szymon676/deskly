package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestCreateBooking(t *testing.T) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForListeningPort("3306/tcp"),
		Env: map[string]string{
			"MYSQL_DATABASE":      "testing",
			"MYSQL_ROOT_PASSWORD": "1234",
			"MYSQL_USER":          "root",
			"MYSQL_PASSWORD":      "1234",
		},
	}

	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		},
	)

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "3306")

	dsn := fmt.Sprintf("root:1234@tcp(127.%s:%s)/testing", host, port)
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
}
