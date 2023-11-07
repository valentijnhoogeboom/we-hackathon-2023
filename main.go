package main

import (
	remote2 "GlobalAPI/controllers/remote"
	scraper2 "GlobalAPI/controllers/scraper"
	user2 "GlobalAPI/controllers/user"
	"GlobalAPI/database"
	"GlobalAPI/metrics"
	"GlobalAPI/middleware"
	"GlobalAPI/types"
	"github.com/fasthttp/session/v2/providers/memory"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/session/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
)

func main() {
	db := createDatabase()
	database.StaticDatabase = &db
	db.Connect()
	db.Migrate()

	metricServer := createInflux()
	metricServer.Start()

	startServer()
}

func createDatabase() database.Database {
	return database.Database{
		Host:     "<redacted>",
		Username: "<redacted>",
		Password: "<redacted>",
		Name:     "<redacted>",
	}
}

func createInflux() *metrics.Metrics {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("<redacted>", "<redacted>")
	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking("<redacted>", "power")
	return &metrics.Metrics{Power: writeAPI}
}

func startServer() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Create a custom validator
	validate := validator.New()

	// Register custom validator with Fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("validator", validate)
		return c.Next()
	})

	api := app.Group("/api")
	api.Post("/guest/register", user2.Register)
	api.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&types.GlobalStatus{Enabled: true})
	})

	remote := api.Group("/remote")
	remote.Use(middleware.HandleRemoteAuthentication)
	remote.Get("/data", remote2.GetData)
	remote.Get("/data/now", remote2.GetDataNow)

	scraper := api.Group("/scraper")
	scraper.Use(middleware.HandleScraperAuthentication)
	scraper.Get("/data", scraper2.GetData)
	scraper.Post("/data", scraper2.PostData)

	user := api.Group("/user")

	user.Get("/meters", user2.GetMeters)

	meter := user.Group("/meter/:meter")
	meter.Use(middleware.HandleMeterInjection)
	meter.Get("/", user2.IndexMeter)
	meter.Get("/data", user2.GetData)
	meter.Post("/data", user2.PostData)

	store, _ := memory.New(memory.Config{}) // Use your chosen store

	newSession := session.New(session.Config{
		Provider: store,
	})

	user.Use(func(c *fiber.Ctx) error {
		c.Locals("session", newSession)
		return c.Next()
	})

	user.Use(middleware.HandleUserAuthorization)

	user.Post("/login", user2.Login)
	user.Post("/logout", user2.Logout)

	log.Fatal(app.Listen("<redacted>"))
}
