package main

import (
	"time"
)

type Booking struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	DeskID    int    `gorm:"column:desk_id"`
	UserID    int    `gorm:"column:user_id"`
	StartTime string `gorm:"column:start_time"`
	EndTime   string `gorm:"column:end_time"`
}

type BookingRequest struct {
	DeskID    int    `gorm:"column:desk_id"`
	UserID    int    `gorm:"column:user_id"`
	StartTime string `gorm:"column:start_time"`
	EndTime   string `gorm:"column:end_time"`
}

func NewBookingFromBookingRequest(req *BookingRequest) *Booking {
	return &Booking{
		DeskID:    req.DeskID,
		UserID:    req.UserID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
}

// check if booking struct is correct, if not return false if yes return true
func VerifyBooking(booking *Booking) bool {
	if booking.DeskID == 0 || booking.UserID == 0 {
		return false
	}

	result := VerifyTime(booking)
	return result
}

func VerifyTime(booking *Booking) bool {
	startTime, err := time.Parse(time.ANSIC, booking.StartTime)
	if err != nil {
		return false
	}

	endTime, err := time.Parse(time.ANSIC, booking.EndTime)
	if err != nil {
		return false
	}

	if endTime.Before(startTime) {
		// End time is before start time
		return false
	}

	return true
}
