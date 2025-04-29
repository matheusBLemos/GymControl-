package usecase

import (
	"errors"
	"fmt"

	"github.com/mathgod152/GymControl/internal/dto"
	"github.com/mathgod152/GymControl/internal/entity"
)

type ExerciceUsecase struct {
	ExerciceInterface entity.ExerciceInterface
}

func (e *ExerciceUsecase) CreateNewExercice(input dto.ExerciceDto, createUser string) (dto.ExerciceDto, error) {
	exercice := entity.ExerciceEntity{
		Name:        input.Name,
		Muscle:      input.Muscle,
		Description: input.Description,
		Equipament:  input.Equipament,
		Video:       input.Video,
		CreateUser:  createUser,
	}
	if err := exercice.IsValidExercice(); err != nil {
		return dto.ExerciceDto{}, errors.New(fmt.Sprintf(" Error to create Exercice: ", err))
	}
	if _, err := e.ExerciceInterface.Create(&exercice); err != nil {
		return dto.ExerciceDto{}, errors.New(fmt.Sprintf("Error to create Exercice", err))
	}
	return dto.ExerciceDto{
		Name:        input.Name,
		Muscle:      input.Muscle,
		Description: input.Description,
		Equipament:  input.Equipament,
		Video:       input.Video,
	}, nil
}

func (e *ExerciceUsecase) FindAllExercices() ([]dto.ReturnExerciceDto, error) {
	exercice, err := e.ExerciceInterface.FindAll(1, 100, "name") // @TODO remover hardcode
	if err != nil {
		return nil, errors.New(fmt.Sprintf(" Error to Get Exercices", err))
	}
	var exerciceFotmated []dto.ReturnExerciceDto
	for _, entity := range exercice {
		exerciceFotmated = append(exerciceFotmated, dto.ReturnExerciceDto{
			Name:        entity.Name,
			Muscle:      entity.Muscle,
			Description: entity.Description,
			Equipament:  entity.Equipament,
			Video:       entity.Video,
		})
	}
	return exerciceFotmated, nil
}

func (e *ExerciceUsecase) FindExerciceById(exerciceId string) (dto.ReturnExerciceDto, error) {
	exercice, err := e.ExerciceInterface.FindByID(exerciceId)
	if err != nil {
		return dto.ReturnExerciceDto{}, errors.New(fmt.Sprintf(" Error to Get Exercices", err))
	}
	return dto.ReturnExerciceDto{
		Name:        exercice.Name,
		Muscle:      exercice.Muscle,
		Description: exercice.Description,
		Equipament:  exercice.Equipament,
		Video:       exercice.Video,
	}, nil
}

func (e *ExerciceUsecase) UpdateeExercice(input dto.ExerciceDto, createUser string) (dto.ExerciceDto, error) {
	exercice := entity.ExerciceEntity{
		Name:        input.Name,
		Muscle:      input.Muscle,
		Description: input.Description,
		Equipament:  input.Equipament,
		Video:       input.Video,
		CreateUser:  createUser,
	}
	if _, err := e.ExerciceInterface.Update(&exercice); err != nil {
		return dto.ExerciceDto{}, errors.New(fmt.Sprintf("Error to update Exercice", err))
	}
	return dto.ExerciceDto{
		Name:   input.Name,
		Muscle: input.Muscle,
	}, nil
}

func (e *ExerciceUsecase) DeleteExercice(exerciceId string) (string, error) {
	_, err := e.ExerciceInterface.Delete(exerciceId)
	if err != nil {
		return "", errors.New(fmt.Sprintf(" Error to delete User", err))
	}
	return "Usuario Deletado com Sucesso", nil
}
