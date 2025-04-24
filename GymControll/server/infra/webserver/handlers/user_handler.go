package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/GymControl/internal/dto"
)

func (_r *Router) CreateUserHandler(c *fiber.Ctx) error {
	var user dto.CreateUserDto
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	_, err := _r.User.CreateNewUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error 500" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": "Atualizado com suceesso",
	})
}