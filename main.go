package main

import (
	"log"

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

func (s *Storage) CreateBooking(b *Booking) error {
	result := s.db.Create(b)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func main() {
	storage, err := NewStorage("username:password@tcp(host:port)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	_ = storage
}
