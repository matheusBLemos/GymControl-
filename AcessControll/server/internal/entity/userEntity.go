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

func NewUsuario(id, nome, senha, email, dataNasc, genero, tipoConta string) (*UserEntity,error){
	usuario := &UserEntity{
		ID: id,
		Name: nome,
		Password: senha,
		Email: email,
		Birthday: dataNasc,
		Gender: genero,
		AcountType: tipoConta,
	}
	err := usuario.IsValidUsuario()
	if err != nil {
		return nil, err
	}
	return usuario, nil
}

func (u *UserEntity) IsValidUsuario() error{
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