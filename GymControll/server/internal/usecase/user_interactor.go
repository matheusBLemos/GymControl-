package usecase

import (
	"errors"
	"fmt"

	"github.com/mathgod152/GymControl/internal/dto"
	"github.com/mathgod152/GymControl/internal/entity"
)

type UserInteractorUsecase struct {
	UserInteractor entity.Logininterface
}

func (u *UserInteractorUsecase) Login(login dto.LoginDto) (string, error) {
	if !u.UserInteractor.VeerifyPassword(login.Email, login.Password) {
		return "", errors.New("Invalid Password")
	}
	token, err := u.UserInteractor.GetUserToken(login.Email, login.Password)
	if err != nil {
		return "", errors.New(fmt.Sprintf(" Error to make Login: ", err))
	}
	return token, nil
}
