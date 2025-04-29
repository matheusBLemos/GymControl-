package dto

type ExerciceDto struct {
	Name        string `json:"name"`
	Muscle      string `json:"muscle"`
	Description string `json:"description"`
	Equipament  string `json:"equipament"`
	Video       string `json:"video_url"`
}

type ReturnExerciceDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Muscle      string `json:"muscle"`
	Description string `json:"description"`
	Equipament  string `json:"equipament"`
	Video       string `json:"video_url"`
}
