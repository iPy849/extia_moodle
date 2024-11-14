package handlers

// import (
// 	"extia/utils"

// 	"github.com/gofiber/fiber/v2"
// )

// type LoginRequest struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8,max=32"`
// }

// func (h *Handler) Login(c *fiber.Ctx) error {
// 	var body = new(LoginRequest)
// 	if err := utils.BindAndValidateStruct(c, body); err != nil {
// 		return err
// 	}

// 	// Login user
// 	user, err := h.userRepository.LoginWithCredentials(body.Email, body.Password)
// 	if err != nil {
// 		return fiber.ErrUnauthorized
// 	}

// 	// Create JWT
// 	token, err := user.EmitJWTToken()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	return c.JSON(fiber.Map{"token": token})
// }

// type RegisterRequest struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8,max=32"`
// }

// func (h *Handler) Register(c *fiber.Ctx) error {
// 	var body = new(RegisterRequest)
// 	if err := utils.BindAndValidateStruct(c, body); err != nil {
// 		return err
// 	}

// 	// Register user
// 	user, _ := h.userRepository.SearchUserByEmail(body.Email)
// 	if user != nil {
// 		return c.Status(fiber.StatusNotAcceptable).SendString("Email already registered")
// 	}

// 	_, err := h.userRepository.RegisterUser(body.Email, body.Password)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotAcceptable).SendString(err.Error())
// 	}

// 	return c.Status(fiber.StatusCreated).Send(nil)
// }

// func (h *Handler) Recover(c *fiber.Ctx) error {
// 	return fiber.ErrNotImplemented
// }
