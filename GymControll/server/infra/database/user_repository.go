package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/mathgod152/GymControl/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

var _ entity.UserInterface = (*UserRepository)(nil)

type UserRepository struct {
	Db *sql.DB
}

func (u *UserRepository) CryptoPassword(password string) (string, error){
	// Define o custo de hashing (14 é seguro, mas pode ser ajustado conforme necessidade)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {
	user.Id = uuid.New().String()
	if err := user.IsValidUser(); err != nil {
		return nil, err
	}
	stmt, err := u.Db.Prepare("INSERT INTO gym_controll_users (name, password, email, birthday, gender, account_type) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(user.Id, user.Name, user.Password, user.Email, user.Gender, user.AcountType)
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

