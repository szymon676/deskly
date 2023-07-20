package main

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
