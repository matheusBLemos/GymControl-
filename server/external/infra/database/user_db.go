package database

import (
	"GusLem/gymControll/internal/entity"
	"database/sql"

	"github.com/google/uuid"
)

var _ entity.UserInterface = (*UserRepository)(nil)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (a *UserRepository) Create(User *entity.UserEntity) (*entity.UserEntity, error) {
	id := uuid.New().String()
	stmt, err := a.Db.Prepare("INSERT INTO User (id, name, password, email, birthday, gender, acount_type) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id, User.Name, User.Password, User.Email, User.Birthday, User.Gender, User.AcountType)
	if err != nil {
		return nil, err
	}
	return &entity.UserEntity{
		ID:         id,
		Name:       User.Name,
		Password:   User.Password,
		Email:      User.Email,
		Birthday:   User.Birthday,
		Gender:     User.Gender,
		AcountType: User.AcountType,
	}, nil
}

func (a *UserRepository) FindAll(page, limit int, sort string) ([]*entity.UserEntity, error) {
	rows, err := a.Db.Query("SELECT id, name, email, birthday, gender, acount_type FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Users := []*entity.UserEntity{}
	for rows.Next() {
		var id, name, email, birthday, gender, acountType string
		if err := rows.Scan(&id, &name, &email, &birthday, &gender, &acountType,); err != nil {
			return nil, err
		}
		User := entity.UserEntity{
			ID:         id,
			Name:       name,
			Email:      email,
			Birthday:   birthday,
			Gender:     gender,
			AcountType: acountType,
		}
		Users = append(Users, &User)
	}
	return Users, nil
}

func (a *UserRepository) FindByID(id string) (*entity.UserEntity, error) {
	rows, err := a.Db.Query("SELECT id, name, email, birthday, gender, acount_type FROM user WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	User := entity.UserEntity{}
	for rows.Next() {
		var id, name, email, birthday, gender, acountType string
		if err := rows.Scan(&id, &name, &email, &birthday, &gender, &acountType,); err != nil {
			return nil, err
		}
		User = entity.UserEntity{
			ID:         id,
			Name:       name,
			Email:      email,
			Birthday:   birthday,
			Gender:     gender,
			AcountType: acountType,
		}
	}
	return &User, nil
}
