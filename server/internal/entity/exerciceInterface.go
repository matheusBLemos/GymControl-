package entity

type ExerciceInterface interface {
	Create(exercise *ExerciceEntity) (*ExerciceEntity, error)
	FindAll(page, limit int, sort string) ([]*ExerciceEntity, error)
	FindByID(id string) (*ExerciceEntity, error)
//	Delete(id string) (string, error)
}
