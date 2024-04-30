package entity

type TrainingInterface interface {
	Create(exercise *TrainingEntity) (TrainingEntity, error)
	FindAll(page, limit int, sort string) ([]*TrainingEntity, error)
	FindByID(id string) (*TrainingEntity, error)
	Delete(id string) (string, error)
}
