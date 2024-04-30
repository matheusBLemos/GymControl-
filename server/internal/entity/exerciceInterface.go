package entity

type ExerciceInterface interface {
	Create(execice *ExerciceEntity) (ExerciceEntity, error)
	FindAll(page, limit int, sort string) ([]*ExerciceEntity, error)
	FindByExerciceID(id string) (*ExerciceEntity, error)
	Delete(id string) (string, error)
}