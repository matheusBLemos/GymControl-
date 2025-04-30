package entity

type WorkoutSectionInterface interface {
	Create(createUser string, workoutSection *WorkoutSectionEntity) (*WorkoutSectionEntity, error)
	CreateExerciseSet(createUser string, exercicesSet *ExerciseSetEntity) (*ExerciseSetEntity, error)
	VerifyExerciseSetOrder(order int32) error
	AssignWorkoutSectionToCustomer(workoutSectionId, personalEmail, customerEmail string) error
}
