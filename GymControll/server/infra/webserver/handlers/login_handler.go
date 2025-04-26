package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/GymControl/internal/dto"
)

func (_r *Router) LoginHandler(c *fiber.Ctx) error {
	var login dto.LoginDto
	if err := c.BodyParser(&login); err != nil {
		fmt.Println("ERRO: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	updatedUser, err := _r.Login.Login(login)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error login: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": updatedUser,
	})
}