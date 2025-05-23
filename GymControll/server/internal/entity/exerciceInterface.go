package entity

type ExerciceInterface interface {
	Create(*ExerciceEntity)(*ExerciceEntity, error)
	FindAll(page, limit int, sort string) ([]*ExerciceEntity, error)
	FindByID(id string) (*ExerciceEntity, error)
	Update(Exercice *ExerciceEntity) (*ExerciceEntity, error)
	Delete(id string) (bool, error)
}

