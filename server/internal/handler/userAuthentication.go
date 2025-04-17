package handler

import (
	"github.com/Hacks-coder/zigi-fashion/internal/database"
	"github.com/Hacks-coder/zigi-fashion/internal/model"
	"github.com/Hacks-coder/zigi-fashion/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var request AuthenticationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := model.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: utils.GeneratePassword(request.Password),
	}
	res := database.DBConnection.Create(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message:": "user created successfully...",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var request AuthenticationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user model.User
	res := database.DBConnection.Where("email = ?", request.Email).First(&user)
	if res.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found.",
		})
	}
	matched, err := utils.ComparePassword(user.PasswordHash, request.Password)
	if err != nil || !matched {
		return c.Status(401).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// token generation
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
