package entity

type WorkoutSectionInterface  interface{
	CreateExerciseSet(*ExerciseSetEntity)(*ExerciseSetEntity)
	Create(*WorkoutSectionEntity)(*WorkoutSectionEntity)
}