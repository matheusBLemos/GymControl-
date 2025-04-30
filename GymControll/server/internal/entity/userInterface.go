package entity

type UserInterface interface {
	CryptoPassword(string) (string, error)
	Create(*User) (*User, error)
	FindAll(page, limit int, sort string) ([]*User, error)
	FindByID(id string) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) (bool, error)
}