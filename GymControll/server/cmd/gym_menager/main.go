package main

import (
	"database/sql"
	"os"
	"os/signal"

	"github.com/mathgod152/GymControl/cmd/dbinit"
	"github.com/mathgod152/GymControl/infra/database"
	"github.com/mathgod152/GymControl/infra/implementations"
	"github.com/mathgod152/GymControl/infra/webserver"
	"github.com/mathgod152/GymControl/internal/entity"
	"github.com/mathgod152/GymControl/internal/usecase"
)

var (
	serverImpl entity.Server
	DB         *sql.DB
	err        error
)

func init() {
	var (
		UserImplementation  = &database.UserRepository{Db: dbinit.DB}
		LoginImplementation = &implementations.LoginImplementations{Db: dbinit.DB}
		ExerciceImplementation = &database.ExerciceRepository{Db: dbinit.DB}

		runUserImplementation = &usecase.UserUsecase{
			UserInterface: UserImplementation, 
		}
		runLoginImplementation = &usecase.UserInteractorUsecase{
			UserInteractor: LoginImplementation,
		}
		runExerciceImplementation = &usecase.ExerciceUsecase{
			ExerciceInterface: ExerciceImplementation,
		}
	)

	serverImpl = &webserver.Server{
		User: runUserImplementation,
		Login: runLoginImplementation,
		Exercice: runExerciceImplementation,
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
