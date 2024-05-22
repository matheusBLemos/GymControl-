package controllers

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func (a *UserController) CreateUser(c *fiber.Ctx) error {
	var user dto.UserDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	newUser, err := a.UserUseCase.CreateNewUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}
	fmt.Println("Response:", newUser)
	return c.JSON(fiber.Map{
		"response": newUser,
	})
}
