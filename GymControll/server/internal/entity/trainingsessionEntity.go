package entity


type ExerciceSerie struct {
	SerieNumber  int8
	NumberAtReps int8
}

type TrainingSessionExercicesToDo struct {
	Exercice   ExerciceEntity
	SeriesToDo []ExerciceSerie
}
type TrainingSessionExercicesComplete struct {
	Exercice     ExerciceEntity
	SeriesReport []ExerciceSerie
}
type TrainingSessionToDoEntity struct {
	ID         string
	Name       string
	Muscles    string
	Exercices  []TrainingSessionExercicesToDo
	CreateUser string
}

type CompletedTrainingSessionEntity struct {
	ID         string
	Name       string
	Muscle     string
	Exercices  []TrainingSessionExercicesComplete
	CreateUser string
}

// func (u *TrainingSessionToDoEntity) IsValidTrainingSessionEntity() error {
// 	if u.ID == "" {
// 		return errors.New("ID is required")
// 	}
// 	if u.Name == "" {
// 		return errors.New("Name is required")
// 	}
// 	if u.Muscle == "" {
// 		return errors.New("Muscle is required")
// 	}
// 	if u.CreateUser == "" {
// 		return errors.New("Email Obrigatorio")
// 	}
// 	return nil
// }

// func (u *CompletedTrainingSessionEntity) IsValidTrainingSessionEntity() error {
// 	if u.ID == "" {
// 		return errors.New("ID is required")
// 	}
// 	if u.Name == "" {
// 		return errors.New("Name is required")
// 	}
// 	if u.Muscle == "" {
// 		return errors.New("Muscle is required")
// 	}
// 	if u.CreateUser == "" {
// 		return errors.New("Email Obrigatorio")
// 	}
// 	return nil
// }
