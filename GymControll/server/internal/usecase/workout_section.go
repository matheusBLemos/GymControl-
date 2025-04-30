package usecase

import (
	"errors"

	"github.com/mathgod152/GymControl/internal/dto"
	"github.com/mathgod152/GymControl/internal/entity"
)

type WorkoutSectionUsecase struct {
	WorkoutSectionUsecase entity.WorkoutSectionInterface
}

func (w *WorkoutSectionUsecase) AssignWorkoutSectionToCustomer(createUser string, inputWorkoutSection dto.WorkoutSectionDto, inputExercicesSet []*dto.ExerciseSetnDto) (dto.WorkoutSectionReturn, error) {
	workoutSection := &entity.WorkoutSectionEntity{
		Name:        inputWorkoutSection.Name,
		Description: inputWorkoutSection.Description,
	}

	var exercicesSet []entity.ExerciseSetEntity
	createWorkoutSection, err := w.WorkoutSectionUsecase.Create(createUser, workoutSection)
	if err != nil {
		return dto.WorkoutSectionReturn{}, err
	}
	for _, exerciceSet := range inputExercicesSet {
		w.WorkoutSectionUsecase.VerifyExerciseSetOrder(exerciceSet.Order) //TODO Refatorar essa validação
		createExerciseSet, err := w.WorkoutSectionUsecase.CreateExerciseSet(
			createUser,
			&entity.ExerciseSetEntity{
				WorkoutSectionId: createWorkoutSection.Id,
				ExerciceId:       exerciceSet.ExerciceId,
				Order:            exerciceSet.Order,
				Sets:             exerciceSet.Sets,
				RepsPerSet:       exerciceSet.RepsPerSet,
			})
		if err != nil {
			return dto.WorkoutSectionReturn{}, errors.New("Error to Create: %v set")
		}
		exercicesSet = append(exercicesSet, *createExerciseSet)
	}
	return dto.WorkoutSectionReturn{
		Name:         createWorkoutSection.Name,
		Description:  createWorkoutSection.Description,
		ExercicseSet: exercicesSet,
	}, nil
}
