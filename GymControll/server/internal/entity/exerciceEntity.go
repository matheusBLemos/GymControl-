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

func (u *ExerciceEntity) IsValidExercice() error {
	if u.Name == "" {
		return errors.New("Name is required")
	}
	if u.Muscle == "" {
		return errors.New("Muscle is required")
	}
	if u.CreateUser == "" {
		return errors.New("Email Obrigatorio")
	}
	return nil
}
