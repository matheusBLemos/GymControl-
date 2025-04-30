package entity

import "errors"

type ExerciseSetEntity struct {
	Id               string
	WorkoutSectionId string
	ExerciceId       string
	Order            int32
	Sets             int32
	RepsPerSet       []string
}

type WorkoutSectionEntity struct {
	Id          string
	Name        string
	Description string
	ExercicesSet []ExerciseSetEntity
}

func (e *ExerciseSetEntity) IsValidExerciseSet() error {
	if e.WorkoutSectionId == "" {
		return errors.New("Name is required")
	}
	if e.ExerciceId == "" {
		return errors.New("Muscle is required")
	}
	if e.Order < 0 {
		return errors.New("Email is required")
	}
	if e.Sets < 0 {
		return errors.New("Muscle is required")
	}
	if len(e.RepsPerSet) < int(e.Sets) {
		return errors.New("Email is required")
	}
	return nil
}