package entity

import "errors"

type ExerciceEntity struct{
	ID string
	Name string
	Desc string
	CreateUser string
	ExerciceName string
}

func CreateExercice(id, name, desc, createUser, exerciceName string) (*ExerciceEntity, error){
	exercice := &ExerciceEntity{
		ID: id,
		Name: name,
		Desc: desc,
		CreateUser: createUser,
		ExerciceName: exerciceName,
	}
	err := exercice.IsValidExercice()
	if err != nil {
		return nil, err
	}
	return exercice, nil
}

func (e *ExerciceEntity) IsValidExercice() error{
	if e.ID == ""{
		return errors.New("Invalid ID")
	}
	if e.Name == ""{
		return errors.New("Invalid ID")
	}
	if e.Desc == ""{
		return errors.New("Invalid ID")
	}
	if e.CreateUser == ""{
		return errors.New("Invalid ID")
	}
	if e.ExerciceName == ""{
		return errors.New("Invalid ID")
	}
	return nil
}