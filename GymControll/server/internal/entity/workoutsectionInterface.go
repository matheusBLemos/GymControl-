package entity

type WorkoutSectionInterface  interface{
	Create(createUser string,workoutSectionId *WorkoutSectionEntity)(*WorkoutSectionEntity)
	CreateExerciseSet(createUser string,exercicesSet *ExerciseSetEntity)(*ExerciseSetEntity)
	AssignWorkoutSectionToCustomer(workoutSectionId, personalEmail, customerEmail string) error
}