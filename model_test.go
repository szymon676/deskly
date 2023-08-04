package main

import (
	"testing"
	"time"
)

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
			StartTime: time.Now().Add(time.Second * 10).Format(time.ANSIC),
			EndTime:   time.Now().Add(time.Second * 15).Format(time.ANSIC),
		}

		ok := VerifyBooking(correctBooking)
		if !ok {
			t.Fatal("should return true, not false")
		}
	})
}
