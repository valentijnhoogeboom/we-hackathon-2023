package middleware

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func HandleMeterInjection(c *fiber.Ctx) error {
	userID := c.Locals("user")
	meterID, err := c.ParamsInt("meter", -1)

	var user models.User
	var meter models.Meter

	database.StaticDatabase.DB.Model(&models.User{}).First(&user, userID)

	if user.ID == 0 {
		return fiber.NewError(http.StatusInternalServerError)
	}

	if err != nil {
		return fiber.NewError(http.StatusNotFound, "Meter was not found.")
	}

	database.StaticDatabase.DB.Model(&models.Meter{}).Where(&models.Meter{UserID: user.ID}).First(&meter, meterID)

	if meter.ID == 0 {
		return fiber.NewError(http.StatusNotFound, "Meter was not found.")
	}

	c.Locals("meter", meter.ID)

	return c.Next()
}
