package entity

import "errors"

type ExerciceEntity struct {
	ID          string
	Name        string
	Muscle      string
	Description string 
	Equipament  string 
	Video       string 
	CreateUser  string
}

func (e *ExerciceEntity) IsValidExercice() error {
	if e.Name == "" {
		return errors.New("Name is required")
	}
	if e.Muscle == "" {
		return errors.New("Muscle is required")
	}
	if e.CreateUser == "" {
		return errors.New("Email is required")
	}
	return nil
}
