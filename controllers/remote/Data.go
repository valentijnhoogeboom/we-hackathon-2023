package remote

import (
	"GlobalAPI/database"
	"GlobalAPI/models"
	"github.com/gofiber/fiber/v2"
)

func GetData(c *fiber.Ctx) error {
	var period = c.Query("period", "hourly")
	var data []models.Data

	switch period {
	case "hourly":
		{
			database.StaticDatabase.DB.Raw(""+
				"SELECT *"+
				"		FROM data"+
				"	WHERE created_at IN ("+
				"	SELECT created_at"+
				"	FROM ("+
				"	SELECT"+
				"created_at,"+
				"	ROW_NUMBER() OVER(PARTITION BY DATE_FORMAT(created_at, '%Y-%m-%d %H') ORDER BY created_at) AS row_num"+
				"FROM data"+
				") AS ranked"+
				"WHERE row_num <= 5"+
				");"+
				"", &data)
			break
		}
	}

	database.StaticDatabase.DB.Model(&models.Data{}).Find(&data)
	return c.JSON(&data)
}

func GetDataNow(c *fiber.Ctx) error {
	var data models.Data
	database.StaticDatabase.DB.Model(&models.Data{}).Last(&data)
	return c.JSON(&data)
}
