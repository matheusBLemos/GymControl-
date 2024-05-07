package main

import (
	"GusLem/gymControll/configs"
	"GusLem/gymControll/external/infra/database"
	"GusLem/gymControll/external/infra/webserver"
	"GusLem/gymControll/internal/web"
)
var (
	err error
)

func init(){
	
}

func main(){
	webserver := webserver.NewWebServer(":3000")
	webClientHandler := web.NewUserHandler(&database.UserRepository{})
	webserver.AddHandler("/order", )
	println("Starting web server on port", configs.Config.WebServerPort)
	go webserver.Start()
}