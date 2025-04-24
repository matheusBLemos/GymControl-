package webserver

import (
	"GusLem/gymControll/cmd/dbinit"
	"GusLem/gymControll/infra/database"
	controllers "GusLem/gymControll/infra/webserver/controller"
	"GusLem/gymControll/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

var TrainingServer *fiber.App
var userController *controllers.UserController

func init() {
	TrainingServer = fiber.New()
	userUseCase := &usecase.UserUseCase{UserInterface: &database.UserRepository{Db: dbinit.Dbinit}}
	userController = &controllers.UserController{UserUseCase: *userUseCase}

	TrainingServer.Get("/users", userController.CreateUser)
	TrainingServer.Post("/send/method", userController.CreateUser)
}