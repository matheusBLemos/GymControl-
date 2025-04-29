package usecase

import (
	"errors"
	"fmt"

	"github.com/mathgod152/GymControl/internal/dto"
	"github.com/mathgod152/GymControl/internal/entity"
)

type UserUsecase struct {
	UserInterface         entity.UserInterface
	UserPersonalInterface entity.UserPersonalInterface
}

func (u *UserUsecase) CreateNewUser(input dto.CreateUserDto) (dto.UserDto, error) {
	password, error := u.UserInterface.CryptoPassword(input.Password)
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
	if err := user.IsValidUser(); err != nil {
		return dto.UserDto{}, err
	}
	if _, err := u.UserInterface.Create(user); err != nil {
		return dto.UserDto{}, errors.New(fmt.Sprintf(" Error to create User", err))
	}
	return dto.UserDto{
		Name:       input.Name,
		Email:      input.Email,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		AcountType: input.AcountType,
	}, nil
}

func (u *UserUsecase) FindAllUsers(page, limit int, sort string) ([]dto.UserDto, error) {
	users, err := u.UserInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(" Error to Get User", err))
	}
	var usersFormated []dto.UserDto

	for _, entity := range users {
		usersFormated = append(usersFormated, dto.UserDto{
			Name:       entity.Name,
			Email:      entity.Email,
			Birthday:   entity.Birthday,
			Gender:     entity.Gender,
			AcountType: entity.AcountType,
		})
	}
	return usersFormated, nil
}

func (u *UserUsecase) FindUserById(userId string) (dto.UserDto, error) {
	users, err := u.UserInterface.FindByID(userId)
	if err != nil {
		return dto.UserDto{}, errors.New(fmt.Sprintf(" Error to create User", err))
	}
	return dto.UserDto{
		Name:       users.Name,
		Email:      users.Email,
		Birthday:   users.Birthday,
		Gender:     users.Gender,
		AcountType: users.AcountType,
	}, nil
}

func (u *UserUsecase) UpdateUser(userID string, input dto.CreateUserDto) (dto.UserDto, error) {
	var err error
	if input.Password != "" {
		input.Password, err = u.UserInterface.CryptoPassword(input.Password)
		if err != nil {
			return dto.UserDto{}, errors.New("Error to Encrypt Password")
		}
	}
	user := &entity.User{
		Id:         userID,
		Name:       input.Name,
		Password:   input.Password,
		Email:      input.Email,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		AcountType: input.AcountType,
	}
	if _, err := u.UserInterface.Update(user); err != nil {
		return dto.UserDto{}, errors.New(fmt.Sprintf(" Error to Update User", err))
	}
	return dto.UserDto{
		Name:       input.Name,
		Email:      input.Email,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		AcountType: input.AcountType,
	}, nil
}

func (u *UserUsecase) DeleteUser(userId string) (string, error) {
	_, err := u.UserInterface.Delete(userId)
	if err != nil {
		return "", errors.New(fmt.Sprintf(" Error to create User", err))
	}
	return "Usuario Deletado com Sucesso", nil
}
