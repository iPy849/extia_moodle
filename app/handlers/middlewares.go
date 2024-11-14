package handlers

import (
	"encoding/json"
	"extia/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	LocalApikeyInfo = "apiKeyInfo"
)

func AuthorizationMiddleware(c *fiber.Ctx) error {
	if c.Cookies(authCookieName) == "" {
		return c.Redirect("/login", http.StatusFound)
	}
	return c.Next()
}

func ApikeyAuthorizationMiddleware(c *fiber.Ctx) error {
	authorizationHeader := c.GetReqHeaders()["Authorization"]
	if len(authorizationHeader) == 0 || authorizationHeader[0] == "" {
		return c.SendStatus(http.StatusUnauthorized)
	}

	domain := new(repository.Domain)
	doesExists := domain.DoesApikeyExists(authorizationHeader[0])
	if !doesExists {
		return c.SendStatus(http.StatusUnauthorized)
	}

	userData, _ := json.Marshal(domain)
	c.Locals(LocalApikeyInfo, string(userData))
	return c.Next()

}
