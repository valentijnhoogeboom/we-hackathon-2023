package user

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"GlobalAPI/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	// Parse and validate user input
	var login types.LoginRequest
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	// Get the custom validator from the context
	validate := c.Locals("validator").(*validator.Validate)

	// Validate the registration request
	if err := validate.Struct(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	var user models.User
	database.StaticDatabase.DB.Model(&models.User{}).First(&user, &models.User{Email: login.Email})

	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Authentication failed"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Authentication failed"})
	}

	// Create a session
	s := c.Locals("session").(*session.Session)
	sess := s.Get(c)
	if err != nil {
		return err
	}

	// Store user information in the session
	sess.Set("user", user.ID)
	if err := sess.Save(); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged in successfully"})
}
