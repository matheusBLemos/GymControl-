package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mathgod152/GymControl/infra/webserver/handlers"
	"github.com/mathgod152/GymControl/infra/webserver/middlewares"
	"github.com/mathgod152/GymControl/internal/entity"
	"github.com/mathgod152/GymControl/internal/usecase"
)

var _ entity.Server = (*Server)(nil)

type Server struct {
	User     *usecase.UserUsecase
	Login    *usecase.UserInteractorUsecase
	Exercice *usecase.ExerciceUsecase
}

func (s *Server) Start(port string) error {
	apiRouter := &handlers.Router{
		User:     s.User,
		Login:    s.Login,
		Exercice: s.Exercice,
	}

	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Apenas para desenvolvimento
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Roteamento da API
	gymcontollrouter := app.Group("gymcontoll/api/v1")

	// Rotas públicas (não protegidas)
	gymcontollrouter.Post("/user", apiRouter.CreateUserHandler)
	gymcontollrouter.Post("/login", apiRouter.LoginHandler)

	// Middleware de autenticação
	gymcontollrouter.Use(middlewares.AuthMiddleware())

	// Rotas protegidas
	// USERS
	gymcontollrouter.Get("/users", apiRouter.FindAllUsersHandler)
	gymcontollrouter.Get("/user/:id", apiRouter.FindUserByIdHandler)
	gymcontollrouter.Put("/user/:id", apiRouter.UpdateUserHandler)
	gymcontollrouter.Put("/deleteuser/:id", apiRouter.DeleteUserHandler)

	// EXERCICE
	gymcontollrouter.Post("/exercice", apiRouter.CreateExerciceHandler)
	gymcontollrouter.Get("/", apiRouter.FindAllExercicesHandler)
	gymcontollrouter.Get("/exercice/:id", apiRouter.FindExerciceByIdHandler)
	gymcontollrouter.Put("/exercice/:id", apiRouter.UpdateExerciceHandler)
	gymcontollrouter.Delete("/exercice/:id", apiRouter.DeleteExerciceHandler)

	return app.Listen(port)
}
