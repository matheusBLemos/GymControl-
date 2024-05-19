package main

import (
	"GusLem/gymControll/configs"
	"GusLem/gymControll/external/infra/webserver"

)
var (
	err error
)

func init(){
	
}

func main(){
	webserver := webserver.NewWebServer(":3000")
	//webserver.AddHandler("/order", )
	println("Starting web server on port", configs.Config.WebServerPort)
	go webserver.Start()
	select{}
}