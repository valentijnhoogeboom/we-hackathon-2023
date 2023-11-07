package user

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetMeters(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user")
	var user models.User
	database.StaticDatabase.DB.Model(&models.User{}).First(&user, userID)

	if user.ID == 0 {
		return fiber.NewError(http.StatusInternalServerError)
	}

	var meters []models.Meter
	database.StaticDatabase.DB.Model(&models.Meter{}).Find(&meters, &models.Meter{UserID: user.ID})

	return ctx.JSON(&meters)
}

func IndexMeter(ctx *fiber.Ctx) error {
	meterID := ctx.Locals("meter")
	var meter models.Meter
	database.StaticDatabase.DB.Model(&models.Meter{}).First(&meter, meterID)

	if meter.ID == 0 {
		return fiber.NewError(http.StatusInternalServerError)
	}

	return ctx.JSON(&meter)
}
