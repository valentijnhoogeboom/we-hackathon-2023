package middleware

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func HandleScraperAuthentication(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	authorization, ok := headers["Authorization"]
	if !ok {
		return fiber.NewError(http.StatusUnauthorized, "You are not authenticated.")
	}
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		return fiber.NewError(http.StatusUnauthorized, "Authorization was invalid.")
	}

	splitToken := strings.Split(authorization, "Bearer ")
	authorization = splitToken[1]

	if authorization == "" {
		return fiber.NewError(http.StatusUnauthorized, "Authorization was invalid.")
	}

	var meter models.Meter
	database.StaticDatabase.DB.Model(&models.Meter{}).First(&meter, &models.Meter{Token: authorization})

	if meter.ID == 0 {
		return fiber.NewError(http.StatusUnauthorized, "Authorization was invalid.")
	}

	c.Locals("meter", fmt.Sprintf("%v", meter.ID))
	return c.Next()
}
