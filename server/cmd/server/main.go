package main

import (
	"log"

	"github.com/Hacks-coder/zigi-fashion/internal/database"
	"github.com/Hacks-coder/zigi-fashion/internal/route"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectToDatabase()

	app := fiber.New()

	route.SetUpRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error in listening to server")
	}
}
