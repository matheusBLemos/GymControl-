package dto

type UserTrainingDto struct {
	ID           string `json:"id"`
	IdUser       string `json:"iduser"`
	IdTraining   string `json:"idtraining"`
	IdCreateUser string `json:"create_user"`
}
