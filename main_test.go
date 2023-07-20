package main

import (
	"testing"
	"time"
)

func TestCreateBooking(t *testing.T) {
	storage, err := NewStorage("root:1234@tcp(127.0.0.1:3306)/testing")
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

func TestVerifyBooking(t *testing.T) {
	// incorrect booking struct number 1
	t.Run("testing incorrect booking struct number 1", func(t *testing.T) {
		incorrectBooking := &Booking{
			UserID:    91321,
			StartTime: time.Now().Add(time.Second * 10).String(),
			EndTime:   time.Now().Add(time.Second * 15).String(),
		}

		ok := VerifyBooking(incorrectBooking)
		if ok {
			t.Fatal("should return false, not true")
		}

	})

	// incorrect booking struct number 2
	t.Run("testing incorrect booking struct number 2", func(t *testing.T) {
		incorrectBooking := &Booking{
			DeskID:  12312,
			UserID:  91321,
			EndTime: time.Now().Add(time.Second * 15).String(),
		}

		ok := VerifyBooking(incorrectBooking)
		if ok {
			t.Fatal("should return false, not true")
		}
	})

	// correct struct
	t.Run("testing correct struct ", func(t *testing.T) {
		correctBooking := &Booking{
			DeskID:    12312,
			UserID:    91321,
			StartTime: time.Now().Add(time.Second * 10).Format(time.RFC3339),
			EndTime:   time.Now().Add(time.Second * 15).Format(time.RFC3339),
		}

		ok := VerifyBooking(correctBooking)
		if !ok {
			t.Fatal("should return true, not false")
		}
	})
}
