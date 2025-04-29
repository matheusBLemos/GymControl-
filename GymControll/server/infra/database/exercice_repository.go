package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/mathgod152/GymControl/internal/entity"
)

var _ entity.ExerciceInterface = (*ExerciceRepository)(nil)

type ExerciceRepository struct {
	Db *sql.DB
}

func (u *ExerciceRepository) Create(exercice *entity.ExerciceEntity) (*entity.ExerciceEntity, error) {
	exercice.ID = uuid.New().String()
	fmt.Println("EXERCICE: ", exercice)
	stmt, err := u.Db.Prepare("INSERT INTO gym_controll_exercices (id, name, muscle, description, equipament, video, created_user) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(exercice.ID, exercice.Name, exercice.Muscle, exercice.Description, exercice.Equipament, exercice.Video, exercice.CreateUser)
	if err != nil {
		return nil, err
	}
	return &entity.ExerciceEntity{
		Name:        exercice.Name,
		Muscle:      exercice.Muscle,
		Description: exercice.Description,
		Equipament:  exercice.Equipament,
		Video:       exercice.Video,
		CreateUser:  exercice.CreateUser,
	}, nil
}

func (u *ExerciceRepository) FindAll(page, limit int, sort string) ([]*entity.ExerciceEntity, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf(`SELECT id, name, muscle, description, equipament, video, created_user
	                      FROM gym_controll_exercices WHERE activate = 0 
	                      ORDER BY %s 
	                      LIMIT $1 OFFSET $2`, sort)

	rows, err := u.Db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercices []*entity.ExerciceEntity
	for rows.Next() {
		var exercice entity.ExerciceEntity
		err := rows.Scan(&exercice.ID, &exercice.Name, &exercice.Muscle, &exercice.Description, &exercice.Equipament, &exercice.Video, &exercice.CreateUser)
		if err != nil {
			return nil, err
		}
		exercices = append(exercices, &exercice)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return exercices, nil
}

func (u *ExerciceRepository) FindByID(id string) (*entity.ExerciceEntity, error) {
	query := `SELECT id, name, muscle, description, equipament, video, created_user 
	          FROM gym_controll_exercices WHERE id = $1 AND activate = 0`

	row := u.Db.QueryRow(query, id)

	var exercice entity.ExerciceEntity
	err := row.Scan(&exercice.ID, &exercice.Name, &exercice.Muscle, &exercice.Description, &exercice.Equipament, &exercice.Video, &exercice.CreateUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &exercice, nil
}

func (u *ExerciceRepository) Update(exercice *entity.ExerciceEntity) (*entity.ExerciceEntity, error) {
	if exercice.ID == "" {
		return nil, errors.New("ID do exercicio é obrigatório para atualização")
	}

	query := "UPDATE gym_controll_exercices SET "
	args := []interface{}{}
	i := 1

	if exercice.Name != "" {
		query += "name = $" + strconv.Itoa(i) + ", "
		args = append(args, exercice.Name)
		i++
	}
	if exercice.Muscle != "" {
		query += "muscle = $" + strconv.Itoa(i) + ", "
		args = append(args, exercice.Muscle)
		i++
	}
	if exercice.Description != "" {
		query += "description = $" + strconv.Itoa(i) + ", "
		args = append(args, exercice.CreateUser)
		i++
	}
	if exercice.Equipament != "" {
		query += "equipament = $" + strconv.Itoa(i) + ", "
		args = append(args, exercice.Muscle)
		i++
	}
	if exercice.Video != "" {
		query += "video = $" + strconv.Itoa(i) + ", "
		args = append(args, exercice.CreateUser)
		i++
	}
	if exercice.CreateUser != "" {
		query += "created_user = $" + strconv.Itoa(i) + ", "
		args = append(args, exercice.Muscle)
		i++
	}

	if len(args) == 0 {
		return nil, errors.New("nenhum campo fornecido para atualizar")
	}

	query = query[:len(query)-2]
	query += " WHERE id = $" + strconv.Itoa(i)
	args = append(args, exercice.ID)

	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return u.FindByID(exercice.ID)
}

func (u *ExerciceRepository) Delete(id string) (bool, error) {
	query := `UPDATE gym_controll_exercices 
	          SET activate = 1 WHERE id = $1`

	result, err := u.Db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
