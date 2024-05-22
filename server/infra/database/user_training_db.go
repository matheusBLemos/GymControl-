package database

import (
	"GusLem/gymControll/internal/entity"
	"database/sql"

	"github.com/google/uuid"
)

type UserTrainingRepository struct {
	Db *sql.DB
}

func NewUserTrainingRepository(db *sql.DB) *UserTrainingRepository {
	return &UserTrainingRepository{Db: db}
}

func (a *UserTrainingRepository) Create(UserTraining *entity.UserTrainingEntity) (*entity.UserTrainingEntity, error) {
	id := uuid.New().String()
	stmt, err := a.Db.Prepare("INSERT INTO user_training (id, fk_user, fk_training, fk_create_user) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id, UserTraining.IdUser, UserTraining.IdTraining, UserTraining.IdCreateUser)
	if err != nil {
		return nil, err
	}
	return &entity.UserTrainingEntity{
		ID:           id,
		IdUser:       UserTraining.IdUser,
		IdTraining:   UserTraining.IdTraining,
		IdCreateUser: UserTraining.IdCreateUser,
	}, nil
}

func (a *UserTrainingRepository) FindAll(page, limit int, sort string) ([]*entity.UserTrainingEntity, error) {
	rows, err := a.Db.Query("SELECT id, fk_user, fk_training, fk_create_user FROM user_training")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	UserTrainings := []*entity.UserTrainingEntity{}
	for rows.Next() {
		var id, idUser, idTraining, idCreateUser string
		if err := rows.Scan(&id, &idUser, &idTraining, &idCreateUser); err != nil {
			return nil, err
		}
		UserTraining := entity.UserTrainingEntity{
			ID:               id,
			IdUser:           idUser,
			IdTraining:       idTraining,
			IdCreateUser:     idCreateUser,
		}
		UserTrainings = append(UserTrainings, &UserTraining)
	}
	return UserTrainings, nil
}

func (a *UserTrainingRepository) FindByID(id string) (*entity.UserTrainingEntity, error) {
	rows, err := a.Db.Query("SELECT id, fk_user, fk_training, fk_create_user FROM user_training WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	UserTraining := entity.UserTrainingEntity{}
	for rows.Next() {
		var id, idUser, idTraining, idCreateUser string
		if err := rows.Scan(&id, &idUser, &idTraining, &idCreateUser); err != nil {
			return nil, err
		}
		UserTraining = entity.UserTrainingEntity{
			ID:               id,
			IdUser:           idUser,
			IdTraining:       idTraining,
			IdCreateUser:     idCreateUser,
		}
	}
	return &UserTraining, nil
}
