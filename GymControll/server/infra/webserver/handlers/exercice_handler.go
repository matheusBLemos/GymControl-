package handlers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/GymControl/internal/dto"
)

func (_r *Router) CreateExerciceHandler(c *fiber.Ctx) error {
	var input dto.ExerciceDto
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Usuário não autenticado"})
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}
	exercice, err := _r.Exercice.CreateNewExercice(input, userID.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(exercice)
}



func (_r *Router) FindAllExercicesHandler(c *fiber.Ctx) error {
	// page, _ := strconv.Atoi(c.Query("page", "1"))
	// limit, _ := strconv.Atoi(c.Query("limit", "10"))
	// sort := c.Query("sort", "name")

	exercices, err := _r.Exercice.FindAllExercices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(exercices)
}

func (_r *Router) FindExerciceByIdHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	exercice, err := _r.Exercice.FindExerciceById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	if exercice.Id == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Exercice not found"})
	}
	return c.Status(fiber.StatusOK).JSON(exercice)
}

func (_r *Router) UpdateExerciceHandler(c *fiber.Ctx) error {
	var input dto.ExerciceDto
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Usuário não autenticado"})
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid payload"})
	}

	exercice, err := _r.Exercice.UpdateeExercice(input, userID.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(exercice)
}

func (_r *Router) DeleteExerciceHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	msg, err := _r.Exercice.DeleteExercice(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": msg})
}
