package entity

import "errors"

type User struct {
	Id         string
	Name       string
	Password   string
	Email      string
	Birthday   string
	Gender     string
	AcountType string
}

func (u *User) IsValidUser() error {
	if u.AcountType != "default_user" &&
		u.AcountType != "personal" &&
		u.AcountType != "nutritionist" &&
		u.AcountType != "coach" &&
		u.AcountType != "gym_manager" {
		return errors.New("Account Type is not accepted")
	}
	if u.Name == "" {
		return errors.New("Name is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.Birthday == "" {
		return errors.New("Birthday is required")
	}
	if u.Gender == "" {
		return errors.New("Gender is required")
	}
	if u.AcountType == "" {
		return errors.New("Acount Type is required")
	}
	return nil
}
