package main

import (
	"GusLem/gymControll/configs"
	"GusLem/gymControll/infra/webserver"

	"github.com/gofiber/fiber/v2"
)
var (
	trainigServer  *fiber.App
)

func init(){
	trainigServer = webserver.TrainingServer
}

func main(){
	trainigServer.Listen(configs.Config.WebServerPort)
	select{}
}