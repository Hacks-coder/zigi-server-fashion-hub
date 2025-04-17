package route

import (
	"github.com/Hacks-coder/zigi-fashion/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	auth := app.Group("/user/auth")
	auth.Post("/register", handler.RegisterUser)
	auth.Post("/login", handler.LoginUser)
}
