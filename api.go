package main

import (
	"github.com/gofiber/fiber"
)

type apiserver struct {
	Storage
}

func (a *apiserver) handleCreateBokking(c *fiber.Ctx) error {
	var req BookingRequest
	err := c.BodyParser(&req)
	if err != nil {
		return nil
	}
	return c.JSON("created booking")
}
