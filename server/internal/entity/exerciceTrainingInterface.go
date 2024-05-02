package entity

type ExerciceTrainingInterface interface {
	Create(exercise *TrainingExerciceEntity) (TrainingExerciceEntity, error)
	FindAll(page, limit int, sort string) ([]*TrainingExerciceEntity, error)
	FindByID(id string) (*TrainingExerciceEntity, error)
	Delete(id string) (string, error)
}
