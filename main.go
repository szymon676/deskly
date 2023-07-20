package main

import (
	"errors"
	"time"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Booking struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
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

type Storage struct {
	db *gorm.DB
}

func NewStorage(dataSourceName string) (*Storage, error) {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Booking{})
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) CreateBooking(booking *Booking) error {
	ok := VerifyBooking(booking)
	if !ok {
		return errors.New("bad booking struct")
	}

	result := s.db.Create(booking)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

type apiserver struct {
	Storage
}

func (a *apiserver) handleCreateBokking(c *fiber.Ctx) error {
	return c.JSON("placeholder")
}

func main() {
	storage, err := NewStorage("root:1234@tcp(127.0.0.1:3306)/main")
	if err != nil {
		panic(err)
	}
	_ = storage
}
