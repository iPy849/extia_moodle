package app

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	BindRoutes(*fiber.App)
}
