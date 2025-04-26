package entity

type Logininterface interface {
	VeerifyPassword(email, password string) bool
	GetUserToken(userName, password string) (string, error)
	Permissions(token string) (string, error)
}
