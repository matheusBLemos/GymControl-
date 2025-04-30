package dto

import "github.com/mathgod152/GymControl/internal/entity"

type WorkoutSectionDto struct {
	Name         string                     `json:"name"`
	Description  string                     `json:"description"`
}

type ExerciseSetnDto struct {
	ExerciceId string   `json:"exercicse_id"`
	Order      int32    `json:"order"`
	Sets       int32    `json:"sets"`
	RepsPerSet []string `json:"reps_per_set"`
}

type WorkoutSectionReturn struct {
	Name         string                     `json:"name"`
	Description  string                     `json:"description"`
	ExercicseSet []entity.ExerciseSetEntity `json:"exercicse_set"`
}
