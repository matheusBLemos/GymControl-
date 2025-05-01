package database

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mathgod152/GymControl/internal/entity"
)

var _ entity.PersonalInteractorInterface = (*PersonalRepository)(nil)

type PersonalRepository struct {
	Db *sql.DB
}

func (p *PersonalRepository) IsTrainer(userId string) (bool, error) {
	query := `SELECT name  
	          FROM gym_controll_users WHERE id = $1 AND activate = 0`

	row := p.Db.QueryRow(query, userId)

	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
		return false, err
	}

	return true, nil
}

func (p *PersonalRepository) AssociateUserWithTrainer(userEmail, personalEmail string) error {
	id := uuid.New().String()
	stmt, err := p.Db.Prepare("INSERT INTO relationate_personal_custumer (id, personal_email, custumer_email) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, userEmail, personalEmail)
	if err != nil {
		return err
	}
	return nil
}

func (u *PersonalRepository) FindUsersByTrainer(personalEmail string) ([]*entity.User, error) {
	query := `
		SELECT u.*
		FROM relationate_personal_custumer rpc
		JOIN gym_controll_users u ON rpc.custumer_email = u.email
		WHERE rpc.personal_email = $1 AND rpc.activate = 1;
	`

	rows, err := u.Db.Query(query, personalEmail)
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

func (u *PersonalRepository) FindTrainersByUser(custumerEmail string) ([]*entity.User, error) {
	query := `
		SELECT u.*
		FROM relationate_personal_custumer rpc
		JOIN gym_controll_users u ON rpc.personal_email = u.email
		WHERE rpc.custumer_email = $1 AND rpc.activate = 1;
	`

	rows, err := u.Db.Query(query, custumerEmail)
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
