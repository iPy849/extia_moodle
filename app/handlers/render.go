package handlers

import (
	"extia/repository"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const authCookieName = "userId"

type credentialsForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type newDomainForm struct {
	Domain string `form:"domain"`
}

type RenderHandler struct{}

func (h *RenderHandler) BindRoutes(app *fiber.App) {
	// Statics
	app.Static("/static", "./statics")

	// Auth
	app.Get("login", renderLogin)
	app.Post("login", handleLogin)
	app.Get("sign-up", renderSignUp)
	app.Post("sign-up", handleSignUp)
	app.Get("logout", handleLogout)

	// Home
	app.Get("", AuthorizationMiddleware, renderHome)

	// Domains
	domainGroup := app.Group("domains", AuthorizationMiddleware)
	domainGroup.Post("", handleNewDomain)
	domainGroup.Get(":id", handleRemoveDomain)
}

func renderLogin(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func handleLogin(c *fiber.Ctx) error {
	formData := new(credentialsForm)
	if err := c.BodyParser(formData); err != nil {
		return RenderLoginWithError(c, err)
	}

	user := new(repository.User)
	err := user.GetUserByEmail(formData.Email)
	if err != nil {
		return RenderLoginWithError(c, err)
	}

	err = user.CompareHash(formData.Password)
	if err != nil {
		return RenderLoginWithError(c, err)
	}

	loggedCookie := new(fiber.Cookie)
	loggedCookie.Name = authCookieName
	loggedCookie.Value = strconv.FormatUint(uint64(user.ID), 10)
	loggedCookie.Expires = time.Now().Add(12 * time.Hour)
	c.Cookie(loggedCookie)

	return c.Redirect("/", http.StatusFound)
}

func renderSignUp(c *fiber.Ctx) error {
	return c.Render("sign_up", nil)
}

func handleSignUp(c *fiber.Ctx) error {
	formData := new(credentialsForm)
	if err := c.BodyParser(formData); err != nil {
		return RenderSignupWithError(c, err)
	}

	user := new(repository.User)
	user.Email = formData.Email
	user.Hash = formData.Password
	if err := user.Create(); err != nil {
		return RenderSignupWithError(c, err)
	}

	return c.Render("login", fiber.Map{
		"feedback": "Se ha creado su cuenta exitosamente",
	})
}

func handleLogout(c *fiber.Ctx) error {
	loggedCookie := new(fiber.Cookie)
	loggedCookie.Name = authCookieName
	loggedCookie.Expires = time.Now()
	c.Cookie(loggedCookie)

	return c.Redirect("/", http.StatusFound)
}

func renderHome(c *fiber.Ctx) error {
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
		"domains": domains,
	})
}

func handleNewDomain(c *fiber.Ctx) error {
	formData := new(newDomainForm)
	if err := c.BodyParser(formData); err != nil {
		return RenderHomeWithError(c, err)
	}

	userId, err := strconv.ParseUint(c.Cookies(authCookieName), 10, 32)
	if err != nil {
		return RenderHomeWithError(c, err)
	}

	domain := new(repository.Domain)
	domain.UserID = uint(userId)
	domain.Name = formData.Domain

	generateApikey := func() string {
		apikey := [64]rune{}
		for i := 0; i < len(apikey); i++ {
			r := rand.UintN(50)
			if r < 25 {
				apikey[i] = rune(65 + (r % 25))
			} else {
				apikey[i] = rune(97 + (r % 25))
			}
		}
		return string(apikey[:])
	}

	key := generateApikey()
	for domain.DoesApikeyExists(key) {
		key = generateApikey()
	}
	domain.Key = key

	if err := domain.Create(); err != nil {
		return RenderHomeWithError(c, err)
	}

	return c.Redirect("/", http.StatusFound)
}

func handleRemoveDomain(c *fiber.Ctx) error {
	domainId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return RenderHomeWithError(c, err)
	}

	domain := new(repository.Domain)
	domain.ID = uint(domainId)
	domain.Delete()

	return c.Redirect("/", http.StatusFound)
}
