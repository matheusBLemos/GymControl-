package database

import (
	"GusLem/gymControll/internal/entity"
	"database/sql"

	"github.com/google/uuid"
)

type ExerciceTrainingRepository struct {
	Db *sql.DB
}

func NewExerciceTrainingRepository(db *sql.DB) *ExerciceTrainingRepository {
	return &ExerciceTrainingRepository{Db: db}
}

func (a *ExerciceTrainingRepository) Create(ExerciceTraining *entity.TrainingExerciceEntity) (*entity.TrainingExerciceEntity, error) {
	id := uuid.New().String()
	stmt, err := a.Db.Prepare("INSERT INTO ExerciceTraining (id, fk_exercice, fk_training, sets, reps, durations_seconds, total_durations_seconds, order) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id, ExerciceTraining.IdExercice, ExerciceTraining.IdTraining, ExerciceTraining.Sets, ExerciceTraining.Reps, ExerciceTraining.DurationSeconds, ExerciceTraining.TotalDurationSeconds, ExerciceTraining.Order)
	if err != nil {
		return nil, err
	}
	return &entity.TrainingExerciceEntity{
		ID:                   id,
		IdExercice:           ExerciceTraining.IdExercice,
		IdTraining:           ExerciceTraining.IdTraining,
		Sets:                 ExerciceTraining.Sets,
		Reps:                 ExerciceTraining.Reps,
		DurationSeconds:      ExerciceTraining.DurationSeconds,
		TotalDurationSeconds: ExerciceTraining.TotalDurationSeconds,
		Order:                ExerciceTraining.Order,
	}, nil
}

func (a *ExerciceTrainingRepository) FindAll(page, limit int, sort string) ([]*entity.TrainingExerciceEntity, error) {
	rows, err := a.Db.Query("SELECT id, fk_exercice, fk_training, sets, reps, durations_seconds, total_durations_seconds, order FROM ExerciceTraining")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ExerciceTrainings := []*entity.TrainingExerciceEntity{}
	for rows.Next() {
		var id, idExercice, idTraining string
		var sets, reps int32
		var durationsSeconds, totalDurationsSeconds, order int64
		if err := rows.Scan(&id, &idExercice, &idTraining, &sets, &reps, &durationsSeconds, &totalDurationsSeconds); err != nil {
			return nil, err
		}
		ExerciceTraining := entity.TrainingExerciceEntity{
			ID:                   id,
			IdExercice:           idExercice,
			IdTraining:           idTraining,
			Sets:                 sets,
			Reps:                 reps,
			DurationSeconds:      durationsSeconds,
			TotalDurationSeconds: totalDurationsSeconds,
			Order:                order,
		}
		ExerciceTrainings = append(ExerciceTrainings, &ExerciceTraining)
	}
	return ExerciceTrainings, nil
}

func (a *ExerciceTrainingRepository) FindByID(id string) (*entity.TrainingExerciceEntity, error) {
	rows, err := a.Db.Query("SELECT id, fk_exercice, fk_training, sets, reps, durations_seconds, total_durations_seconds, order FROM ExerciceTraining WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ExerciceTraining := entity.TrainingExerciceEntity{}
	for rows.Next() {
		var id, idExercice, idTraining string
		var sets, reps int32
		var durationsSeconds, totalDurationsSeconds, order int64
		if err := rows.Scan(&id, &idExercice, &idTraining, &sets, &reps, &durationsSeconds, &totalDurationsSeconds); err != nil {
			return nil, err
		}
		ExerciceTraining = entity.TrainingExerciceEntity{
			ID:                   id,
			IdExercice:           idExercice,
			IdTraining:           idTraining,
			Sets:                 sets,
			Reps:                 reps,
			DurationSeconds:      durationsSeconds,
			TotalDurationSeconds: totalDurationsSeconds,
			Order:                order,
		}
	}
	return &ExerciceTraining, nil
}
