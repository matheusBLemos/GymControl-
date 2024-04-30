package entity

type TrainingInterface interface{
	Create(execice *TrainingEntity) (TrainingEntity, error)
	FindAll(page, limit int, sort string) ([]*TrainingEntity, error)
	FindByTrainingID(id string) (*TrainingEntity, error)
	Delete(id string) (string, error)
}