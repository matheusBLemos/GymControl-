package implementations

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/mathgod152/GymControl/configs"
	"github.com/mathgod152/GymControl/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

var _ entity.Logininterface = (*LoginImplementations)(nil)

type LoginImplementations struct {
	Db *sql.DB
}

// CompareHashAndPassword(hashedPassword, password []byte) error
func (u *LoginImplementations) VeerifyPassword(email, password string) bool {
	query := `SELECT password FROM gym_controll_users WHERE email = $1 AND activate = 0`

	row := u.Db.QueryRow(query, email)

	var hashpassword []byte
	err := row.Scan(&hashpassword) 
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashpassword, []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (u *LoginImplementations) GetUserToken(email, password string) (string, error) {
	query := `SELECT id, name FROM gym_controll_users WHERE email = $1 AND activate = 0`

	row := u.Db.QueryRow(query, email)

	var id string
	var name string

	err := row.Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("User not found")
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"name":    name,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := []byte(configs.Config.DBName)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *LoginImplementations) Permissions(token string) (string, error) {
	return "token", nil
}
