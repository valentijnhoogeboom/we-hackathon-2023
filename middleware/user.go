package middleware

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	session2 "github.com/gofiber/session/v2"
	"net/http"
)

func HandleUserAuthorization(c *fiber.Ctx) error {
	ses := c.Locals("session").(*session2.Session)
	session := ses.Get(c)
	userID := session.Get("user")

	if userID == 0 {
		return fiber.NewError(http.StatusUnauthorized, "Your session is no longer active, please sign-in again.")
	}

	var user models.User
	fmt.Printf("\n%v", database.StaticDatabase.DB)
	database.StaticDatabase.DB.First(&user, userID)

	if user.ID == 0 {
		return fiber.NewError(http.StatusUnauthorized, "Your session is no longer active, please sign-in again.")
	}

	// Use the session as needed within route handlers
	c.Locals("user", user.ID)

	return c.Next()
}
