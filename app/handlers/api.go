package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type APIHandler struct{}

func (a *APIHandler) BindRoutes(app *fiber.App) {
	// API
	api := app.Group("/api", ApikeyAuthorizationMiddleware)
	api.Get("", handleAPIIndex)
	api.All("alive", func(c *fiber.Ctx) error { return c.SendStatus(http.StatusOK) })
}

type ResponseAPIIndex struct {
	Name      string
	CreatedAt time.Time
}

func handleAPIIndex(c *fiber.Ctx) error {
	domain := new(ResponseAPIIndex)
	json.Unmarshal(
		[]byte(c.Locals(LocalApikeyInfo).(string)),
		domain,
	)
	return c.JSON(domain)
}
