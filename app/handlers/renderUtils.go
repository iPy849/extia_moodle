package handlers

import (
	"extia/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RenderLoginWithError(c *fiber.Ctx, err error) error {
	return c.Render("login", fiber.Map{
		"error": err.Error(),
	})
}

func RenderSignupWithError(c *fiber.Ctx, err error) error {
	return c.Render("sign_up", fiber.Map{
		"error": err.Error(),
	})
}

func RenderHomeWithError(c *fiber.Ctx, e error) error {
	userId, err := strconv.ParseUint(c.Cookies(authCookieName), 10, 32)
	if err != nil {
		return RenderHomeWithError(c, err)
	}

	domain := new(repository.Domain)
	domains, err := domain.GetAllKeysByUser(uint(userId))
	if err != nil {
		return RenderHomeWithError(c, err)
	}

	return c.Render("home", fiber.Map{
		"error":   e.Error(),
		"domains": domains,
	})
}
