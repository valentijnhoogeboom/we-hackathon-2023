package user

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"GlobalAPI/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	// Parse and validate user input
	var newUser types.RegisterUser
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	// Get the custom validator from the context
	validate := c.Locals("validator").(*validator.Validate)

	// Validate the registration request
	if err := validate.Struct(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	// Hash the user's password (use a password hashing library like bcrypt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	// Save the user to the database
	newUser.Password = string(hashedPassword)
	database.StaticDatabase.DB.Create(&models.User{Username: newUser.Username, Email: newUser.Email, Password: string(hashedPassword)})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}
