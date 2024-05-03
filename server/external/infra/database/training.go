package database

import (
	"GusLem/gymControll/internal/entity"
	"database/sql"

	"github.com/google/uuid"
)

type TrainingRepository struct {
	Db *sql.DB
}

func NewTrainingRepository(db *sql.DB) *TrainingRepository {
	return &TrainingRepository{Db: db}
}

func (a *TrainingRepository) Create(Training *entity.TrainingEntity) (*entity.TrainingEntity, error) {
	id := uuid.New().String()
	stmt, err := a.Db.Prepare("INSERT INTO Training (id, fk_user, fk_create_user, fk_training, desc) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id, Training.IdUser, Training.IdCreateUser, Training.IdTraining, Training.Desc)
	if err != nil {
		return nil, err
	}
	return &entity.TrainingEntity{
		ID:           id,
		IdUser:       Training.IdUser,
		IdCreateUser: Training.IdCreateUser,
		IdTraining:   Training.IdTraining,
		Desc:         Training.Desc,
	}, nil
}

func (a *TrainingRepository) FindAll(page, limit int, sort string) ([]*entity.TrainingEntity, error) {
	rows, err := a.Db.Query("SELECT id, fk_user, fk_create_user, fk_training, desc FROM Training")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Trainings := []*entity.TrainingEntity{}
	for rows.Next() {
		var id, idUser, idCreateUser, idTraining, desc string
		if err := rows.Scan(&id, &idUser, &idCreateUser, &idTraining, &desc); err != nil {
			return nil, err
		}
		Training := entity.TrainingEntity{
			ID:           id,
			IdUser:       idUser,
			IdCreateUser: idCreateUser,
			IdTraining:   idTraining,
			Desc:         desc,
		}
		Trainings = append(Trainings, &Training)
	}
	return Trainings, nil
}

func (a *TrainingRepository) FindByID(id string) (*entity.TrainingEntity, error) {
	rows, err := a.Db.Query("SELECT id, fk_user, fk_create_user, fk_training, desc FROM Training WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Training := entity.TrainingEntity{}
	for rows.Next() {
		var id, idUser, idCreateUser, idTraining, desc string
		if err := rows.Scan(&id, &idUser, &idCreateUser, &idTraining, &desc); err != nil {
			return nil, err
		}
		Training = entity.TrainingEntity{
			ID:           id,
			IdUser:       idUser,
			IdCreateUser: idCreateUser,
			IdTraining:   idTraining,
			Desc:         desc,
		}
	}
	return &Training, nil
}
