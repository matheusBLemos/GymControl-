package entity

import "errors"

type UserEntity struct {
	ID string
	Name string
	Password string
	Email string
	Birthday string
	Gender string
	AcountType string	
}

func (u *UserEntity) IsValidUser() error{
	if u.ID == ""{
		return errors.New("ID Obrigatorio")
	}
	if u.Name == ""{
		return errors.New("Nome Obrigatorio")
	}
	if u.Password == ""{
		return errors.New("Senha Obrigatorio")
	}
	if u.Email == ""{
		return errors.New("Email Obrigatorio")
	}
	if u.Birthday == ""{
		return errors.New("DataNasc Obrigatorio")
	}
	if u.Gender == ""{
		return errors.New("Genero Obrigatorio")
	}
	if u.AcountType == ""{
		return errors.New("TipoConta Obrigatorio")
	}
	return nil
}