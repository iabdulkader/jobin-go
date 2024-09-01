package main

import (
	"github.com/gofiber/fiber/v2"

	"jobin/handlers"

	"jobin/db"
	"log"
)



func main() {
	app := fiber.New()
	app.Static("/", "./static")

	if err := db.Initialize(); err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

	handlers.SetupRoutes(app)


	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
