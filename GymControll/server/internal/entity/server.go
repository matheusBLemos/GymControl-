package entity

type Server interface {
	Start(port string) error
}