package entity

type UserInterface interface {
	Create(user *UserEntity) (*UserEntity, error)
	FindAll(page, limit int, sort string) ([]*UserEntity, error)
	FindByID(id string) (*UserEntity, error)
	Update(user *UserEntity) (UserEntity, error)
	Delete(id string) (string, error)
}