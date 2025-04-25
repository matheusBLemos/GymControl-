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

func (_r *Router) FindAllUsersHandler(c *fiber.Ctx) error {
	page, limit, sort := 1, 10, "name" 
	if err := c.QueryParser(&fiber.Map{"page": &page, "limit": &limit, "sort": &sort}); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query parameters",
		})
	}

	users, err := _r.User.FindAllUsers(page, limit, sort)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching users: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}

func (_r *Router) FindUserByIdHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User ID is required",
		})
	}

	user, err := _r.User.FindUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching user: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func (_r *Router) UpdateUserHandler(c *fiber.Ctx) error {
	var user dto.CreateUserDto
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User ID is required",
		})
	}

	updatedUser, err := _r.User.UpdateUser(userId,user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating user: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": updatedUser,
	})
}

func (_r *Router) DeleteUserHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User ID is required",
		})
	}

	message, err := _r.User.DeleteUser(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting user: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}

