package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/mathgod152/GymControl/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

var _ entity.UserInterface = (*UserRepository)(nil)

type UserRepository struct {
	Db *sql.DB
}

func (u *UserRepository) CryptoPassword(password string) (string, error) {
	// Define o custo de hashing (14 é seguro, mas pode ser ajustado conforme necessidade)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {
	user.Id = uuid.New().String()
	stmt, err := u.Db.Prepare("INSERT INTO gym_controll_users (id, name, password, email, birthday, gender, account_type) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(user.Id, user.Name, user.Password, user.Email, user.Birthday, user.Gender, user.AcountType)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Name:       user.Name,
		Email:      user.Email,
		Birthday:   user.Birthday,
		Gender:     user.Gender,
		AcountType: user.AcountType,
	}, nil
}

func (u *UserRepository) FindAll(page, limit int, sort string) ([]*entity.User, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf(`SELECT id, name, password, email, birthday, gender, account_type 
	                      FROM gym_controll_users WHERE activate = 0 
	                      ORDER BY %s 
	                      LIMIT $1 OFFSET $2`, sort)

	rows, err := u.Db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Birthday, &user.Gender, &user.AcountType)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) FindByID(id string) (*entity.User, error) {
	query := `SELECT id, name, email, birthday, gender, account_type 
	          FROM gym_controll_users WHERE id = $1 AND activate = 0`

	row := u.Db.QueryRow(query, id)

	var user entity.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Gender, &user.AcountType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err 
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Update(user *entity.User) (*entity.User, error) {
	if user.Id == "" {
		return nil, errors.New("ID do usuário é obrigatório para atualização")
	}

	query := "UPDATE gym_controll_users SET "
	args := []interface{}{}
	i := 1

	if user.Name != "" {
		query += "name = $" + strconv.Itoa(i) + ", "
		args = append(args, user.Name)
		i++
	}
	if user.Password != "" {
		query += "password = $" + strconv.Itoa(i) + ", "
		args = append(args, user.Password)
		i++
	}
	if user.Email != "" {
		query += "email = $" + strconv.Itoa(i) + ", "
		args = append(args, user.Email)
		i++
	}
	if user.Birthday != "" {
		query += "birthday = $" + strconv.Itoa(i) + ", "
		args = append(args, user.Birthday)
		i++
	}
	if user.Gender != "" {
		query += "gender = $" + strconv.Itoa(i) + ", "
		args = append(args, user.Gender)
		i++
	}
	if user.AcountType != "" {
		query += "account_type = $" + strconv.Itoa(i) + ", "
		args = append(args, user.AcountType)
		i++
	}

	if len(args) == 0 {
		return nil, errors.New("nenhum campo fornecido para atualizar")
	}

	query = query[:len(query)-2] 
	query += " WHERE id = $" + strconv.Itoa(i)
	args = append(args, user.Id)

	stmt, err := u.Db.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return u.FindByID(user.Id)
}

func (u *UserRepository) Delete(id string) (bool, error) {
	query := `UPDATE gym_controll_users 
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
