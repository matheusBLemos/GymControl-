package entity

import "errors"

type UserEntity struct {
	Id string
	Name string
	Password string
	Email string
	Birthday string
	Gender string
	AcountType string	
}

func (u *UserEntity) IsValidUser() error{
	if u.Id == ""{
		return errors.New("ID is required")
	}
	if u.Name == ""{
		return errors.New("Name is required")
	}
	if u.Password == ""{
		return errors.New("Password is required")
	}
	if u.Email == ""{
		return errors.New("Email is required")
	}
	if u.Birthday == ""{
		return errors.New("Birthday is required")
	}
	if u.Gender == ""{
		return errors.New("Gender is required")
	}
	if u.AcountType == ""{
		return errors.New("Acount Type is required")
	}
	return nil
}