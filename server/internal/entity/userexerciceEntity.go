package entity

import "errors"

type UserTrainingEntity struct {
	ID           string
	IdUser       string
	IdTraining   string
	IdCreateUser string
}

func NewUserTraining(id, idUser, idTraining, idCreateUser string) (*UserTrainingEntity, error) {
	userTraining := &UserTrainingEntity{
		ID:           id,
		IdUser:       idUser,
		IdTraining:   idTraining,
		IdCreateUser: idCreateUser,
	}
	err := userTraining.IsValidUserTraining()
	if err != nil {
		return nil, err
	}
	return userTraining, nil
}

func (u *UserTrainingEntity) IsValidUserTraining() error {
	if u.ID == "" {
		return errors.New("Invalid ID")
	}
	if u.IdUser == "" {
		return errors.New("Invalid IdUser")
	}
	if u.IdTraining == "" {
		return errors.New("Invalid IdTraining")
	}
	if u.IdCreateUser == "" {
		return errors.New("Invalid IdCreateUser")
	}
	return nil
}
