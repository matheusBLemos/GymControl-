package dto

type ExerciceDto struct {
	Name   string `json:"name"`
	Muscle string `json:"muscle"`
}

type ReturnExerciceDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Muscle string `json:"muscle"`
}
