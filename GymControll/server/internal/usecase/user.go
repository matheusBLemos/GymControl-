package usecase

import "github.com/mathgod152/GymControl/internal/entity"

type UserUsecase struct {
	UserInterface entity.UserEntity
}

func (U *UserUsecase) CreateNewUser(input dto.User) (dto.User)