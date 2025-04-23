package usecase

import (
	"errors"

	"github.com/mathgod152/GymControl/internal/dto"
	"github.com/mathgod152/GymControl/internal/entity"
)

type UserUsecase struct {
	UserInterface entity.UserInterface
	UserPersonalInterface entity.UserPersonalInterface
}

func (u *UserUsecase) CreateNewUser(input dto.CreateUserDto) (dto.UserDto, error) {
	password, error := u.UserOperations.CryptoPassword(input.Password)
	if error != nil {
		return dto.UserDto{}, errors.New("Error to Encrypt Password")
	}
	user := &entity.User{
		Name:       input.Name,
		Password:   password,
		Email:      input.Email,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		AcountType: input.AcountType,
	}
	err := user.IsValidUser(){
		return dto.UserDto{}, err
	}
	if _, err := u.UserOperations.CreateUser(&user); err != nil {
		return dto.UserDto{}, errors.New("Error to create User")
	}
	return dto.UserDto{
		Name:       input.Name,
		Email:      input.Email,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		AcountType: input.AcountType,
	}, nil
}
