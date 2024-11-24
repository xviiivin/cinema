package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/xviiivin/cinema/database"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
	app.Listen(":3000")
}
