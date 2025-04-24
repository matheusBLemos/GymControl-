package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/mathgod152/GymControl/infra/webserver/handlers"
	"github.com/mathgod152/GymControl/internal/entity"
	"github.com/mathgod152/GymControl/internal/usecase"
)

var _ entity.Server = (*Server)(nil)

type Server struct {
	User *usecase.UserUsecase
}

func (s *Server) Start(port string) error {
	apiRouter := &handlers.Router{
		User: s.User,
	}

	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // APENAS PARA DEV
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Rote API
	gymcontollrouter := app.Group("gymcontoll/api/v1")


	//USERS
	gymcontollrouter.Post("gymcontoll/user", apiRouter.CreateUserHandler)

	return app.Listen(port)
}
