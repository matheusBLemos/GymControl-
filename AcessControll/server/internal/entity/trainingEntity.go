package entity

import "errors"

type TrainingEntity struct{
	ID string
	IdUser string
	IdCreateUser string
	IdTraining string
	Desc string
}

func NewTraining(id, idUser, idCreateUser, idTraining, desc string) (*TrainingEntity, error){
	training := &TrainingEntity{
		ID: id,
		IdUser: idUser,
		IdCreateUser: idCreateUser,
		IdTraining: idTraining,
		Desc: desc,
	}
	err := training.IsValidTraining()
	if err != nil{
		return nil, err
	}
	return training, nil
}

func (t *TrainingEntity) IsValidTraining() error{
	if t.ID == ""{
		return errors.New("Invalid ID")
	}
	if t.IdUser == ""{
		return errors.New("Invalid IdUser")
	}
	if t.IdCreateUser == ""{
		return errors.New("Invalid IdCreateUser")
	}
	if t.IdTraining == ""{
		return errors.New("Invalid IdTraining")
	}
	if t.Desc == ""{
		return errors.New("Invalid Desc")
	}
	return nil
}