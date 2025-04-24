package entity

import "errors"

type TrainingExerciceEntity struct {
	ID         string
	IdExercice string
	IdTraining string
	Sets int32
	Reps int32
	DurationSeconds int64
	TotalDurationSeconds int64
	Order int64
}

func CreateTrainingExercice(id, idExercice, idTraining string, sets, reps int32, durationSeconds, totalDurationsSeconds, Order int64) (*TrainingExerciceEntity, error){
	trainingExercice := &TrainingExerciceEntity{
		ID: id,
		IdExercice: idExercice,
		IdTraining: idTraining,
		Sets: sets,
		Reps: reps,
		DurationSeconds: durationSeconds,
		TotalDurationSeconds: totalDurationsSeconds,
		Order: Order,
	}
	err := trainingExercice.IsValidTrainingExercice()
	if err != nil{
		return nil, err
	}
	return trainingExercice, nil
}

func (t *TrainingExerciceEntity) IsValidTrainingExercice() error{
	if t.ID == ""{
		return errors.New("Invalid ID")
	}
	if t.IdExercice == ""{
		return errors.New("Invalid ID")
	}
	if t.IdTraining == ""{
		return errors.New("Invalid ID")
	}
	return nil
}
