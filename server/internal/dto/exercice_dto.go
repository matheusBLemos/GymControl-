package dto

type ExerciceDto struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	CreateUser   string `json:"create_user"`
	ExerciceName string `json:"exercice_name"`
}
