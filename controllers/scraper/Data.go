package scraper

import (
	"GlobalAPI/database"
	"GlobalAPI/metrics"
	"GlobalAPI/models"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"net/http"
	"strconv"
)

func GetData(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Locals("meter").(string))
	if err != nil {
		fmt.Printf("\nAn error occured while loading meter: %v", err)
		return fiber.NewError(http.StatusInternalServerError, "An internal server error occured while processing this request.")
	}
	var meter models.Meter
	database.StaticDatabase.DB.Model(&models.Meter{}).First(&meter, id)

	if meter.ID == 0 {
		fmt.Printf("\nAn error occured while loading meter: Meter was not found (%v).", id)
		return fiber.NewError(http.StatusInternalServerError, "Could not load meter data")
	}

	var data models.Data
	database.StaticDatabase.DB.Model(&models.Data{}).Find(&data, &models.Data{MeterID: meter.ID})
	return c.JSON(&data)
}

func PostData(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Locals("meter").(string))
	if err != nil {
		fmt.Printf("\nAn error occured while loading meter: %v", err)
		return fiber.NewError(http.StatusInternalServerError, "An internal server error occured while processing this request.")
	}
	var meter models.Meter
	database.StaticDatabase.DB.Model(&models.Meter{}).First(&meter, id)

	if meter.ID == 0 {
		fmt.Printf("\nAn error occured while loading meter: Meter was not found (%v).", id)
		return fiber.NewError(http.StatusInternalServerError, "Could not load meter data")
	}

	var dataOld []models.Data
	var data []models.Data

	// Parse the JSON request into the RequestData struct
	if err := c.BodyParser(&dataOld); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	for _, datum := range dataOld {
		datum.MeterID = meter.ID
		datum.Meter = meter
		data = append(data, datum)

		p := influxdb2.NewPoint(fmt.Sprintf("%v", datum.MeterID),
			map[string]string{"unit": "watt"},
			map[string]interface{}{"current": (datum.ActivePowerW) * (10000.0) / 3600.0},
			datum.CreatedAt)
		err = metrics.GlobalMetrics.WritePoint(context.Background(), p)
	}

	batchSize := 50 // Set your desired batch size

	for i := 0; i < len(data); i += batchSize {
		// Calculate the end index of the batch
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}

		// Get a slice of the current batch
		batch := data[i:end]

		// Insert the current batch into the database
		if err := database.StaticDatabase.DB.Create(&batch).Error; err != nil {
			// Handle the error as needed
			fmt.Printf("Error inserting batch: %v", err)
		}
	}

	database.StaticDatabase.DB.Model(&models.Meter{}).Where(meter.ID).First(&meter)

	return c.JSON(&meter)
}
