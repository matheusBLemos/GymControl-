package entity

type PersonalInteractorInterface interface {
	IsTrainer(userId string) (bool, error)
	AssociateUserWithTrainer(userEmail, personalEmail string) error
	FindUsersByTrainer(personalEmail string) ([]*User, error)
	FindTrainersByUser(userId string) ([]*User, error)
}
