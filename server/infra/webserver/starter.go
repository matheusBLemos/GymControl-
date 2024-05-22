package webserver

import (
	"GusLem/gymControll/infra/database"
	controllers "GusLem/gymControll/infra/webserver/controller"
	"GusLem/gymControll/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

var SDKServer *fiber.App
var UserController *controllers.UserController

func init() {
	SDKServer = fiber.New()
	userUseCase := &usecase.UserUseCase{UserInterface: &database.UserRepository{}}
	apiController = &controllers.ApiController{
		QueryUseCase:  *queryUseCase,
		MethodUseCase: *methodUseCase,
	}

	SDKServer.Get("/request/:customer/:application/:messageTitle/:requestType", apiController.SendRequest)
	SDKServer.Post("/send/method", apiController.SendMethod)
}