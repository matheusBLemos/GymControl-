package entity


type UserTrainingInterface interface {
	Create(userTraining *UserTrainingEntity) (UserTrainingEntity, error)
	FindAll(page, limit int, sort string) ([]*UserTrainingEntity, error)
	FindByUserTrainingID(id string) (*UserTrainingEntity, error)
	Delete(id string) (string, error)
}