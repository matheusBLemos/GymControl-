package main

import (
	"database/sql"
	"os"
	"os/signal"

	"github.com/mathgod152/GymControl/cmd/dbinit"
	"github.com/mathgod152/GymControl/infra/database"
	"github.com/mathgod152/GymControl/infra/webserver"
	"github.com/mathgod152/GymControl/internal/entity"
	"github.com/mathgod152/GymControl/internal/usecase"
)

var (
	serverImpl entity.Server
	DB *sql.DB
	err error
)

func init(){
	var(
		UserImplementation = &database.UserRepository{ Db: dbinit.DB}

		runUserValidator = &usecase.UserUsecase{
			UserInterface: UserImplementation, // Injetando a implementação 
		}
	)

	serverImpl = &webserver.Server{
		User: runUserValidator,
	}
}


func main() {
	go serverImpl.Start(":5000")

	listen()
}

func listen() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	os.Exit(130)
}