package entity

type UserInterface interface {
	CreateUser(user *UserEntity) (UserEntity, error)
	FindAllUsers(page, limit int, sort string) ([]*UserEntity, error)
	FindUserByID(id string) (*UserEntity, error)
	Delete(id string) (string, error)
}