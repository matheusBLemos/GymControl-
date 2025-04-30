package entity

type PersonalInteractorInterface interface {
	Relationate(userEmail, personalEmail string) (*User, error)
	FindUserByTrainer(personalId string) ([]*User, error)
	FindTrainerByUser(userId string) ([]*User, error)
	SetWorkoutSectionToCustumer(workoutSectionId, userId string)(bool)
}