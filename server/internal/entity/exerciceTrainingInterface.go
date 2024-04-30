package entity

type TrainingExerciceInterface interface{
	Create(execice *TrainingExerciceEntity) (TrainingExerciceEntity, error)
	FindAll(page, limit int, sort string) ([]*TrainingExerciceEntity, error)
	FindByTrainingExerciceID(id string) (*TrainingExerciceEntity, error)
	Delete(id string) (string, error)
}