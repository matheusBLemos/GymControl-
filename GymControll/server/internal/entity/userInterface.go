package entity

type UserInterface interface {
	CreateUser(*UserEntity)(*UserEntity, error)
	Relationate(userEmail string) (*UserEntity, error)
	FindAll(page, limit int, sort string) ([]*UserEntity, error)
	FindByID(id string) (*UserEntity, error)
	FindUserByTrainer(personalId string) ([]*UserEntity, error)
	FindTrainerByUser(userId string) ([]*UserEntity, error)
	Update(user *UserEntity) (UserEntity, error)
	Delete(id string) (string, error)
}