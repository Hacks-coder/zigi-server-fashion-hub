package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("I am working...")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error in listening to server")
	}
}
