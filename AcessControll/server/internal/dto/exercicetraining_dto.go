package dto

type TrainingExerciceDto struct {
	ID                   string `json:"id"`
	IdExercice           string `json:"id_exercice"`
	IdTraining           string `json:"id_training"`
	Sets                 int32 `json:"sets"`
	Reps                 int32 `json:"reps"`
	DurationSeconds      int64 `json:"duration_seconds"`
	TotalDurationSeconds int64 `json:"total_duration_seconds"`
	Order                int64 `json:"order"`
}
