package main

import "time"

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

// check if booking struct is correct, if not return false if yes return true
func VerifyBooking(booking *Booking) bool {
	if booking.DeskID == 0 || booking.UserID == 0 {
		return false
	}

	startTime, err := time.Parse(time.RFC3339, booking.StartTime)
	if err != nil {
		// Invalid start time format
		return false
	}

	endTime, err := time.Parse(time.RFC3339, booking.EndTime)
	if err != nil {
		// Invalid end time format
		return false
	}

	if endTime.Before(startTime) {
		// End time is before start time
		return false
	}

	return true
}
