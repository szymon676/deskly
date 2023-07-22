package main

import (
	"github.com/gofiber/fiber/v2"
)

type apiserver struct {
	storage *Storage
}

func (a *apiserver) handleCreateBooking(c *fiber.Ctx) error {
	var req *BookingRequest
	err := c.BodyParser(&req)
	if err != nil {
		return nil
	}

	booking := NewBookingFromBookingRequest(req)

	err = a.storage.CreateBooking(booking)
	if err != nil {
		return err
	}

	return c.JSON("created booking")
}

func (a *apiserver) SetupServer() *fiber.App {
	app := fiber.New(fiber.Config{DisableStartupMessage: false})

	app.Post("/booking", a.handleCreateBooking)

	return app
}
