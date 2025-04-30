package entity

type PersonalInteractor interface {
	Relationate(userEmail, personalEmail string) (*User, error)
	FindUserByTrainer(personalId string) ([]*User, error)
	FindTrainerByUser(userId string) ([]*User, error)
}