package app

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var errCannotCompleteRouteBinding = errors.New("cannot complete routes binding")

func BindHandlersToApp(app *fiber.App, handlers ...Handler) error {
	for _, handler := range handlers {
		handler.BindRoutes(app)
	}
	return nil
}
