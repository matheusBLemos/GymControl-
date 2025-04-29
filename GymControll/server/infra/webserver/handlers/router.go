package handlers

import "github.com/mathgod152/GymControl/internal/usecase"

type Router struct {
	User *usecase.UserUsecase
	Login *usecase.UserInteractorUsecase
	Exercice *usecase.ExerciceUsecase
}