package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

func Logout(c *fiber.Ctx) error {
	s := c.Locals("session").(*session.Session)

	// Retrieve the session
	sess := s.Get(c)

	// Destroy the session
	err := sess.Destroy()
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not logout."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
}
