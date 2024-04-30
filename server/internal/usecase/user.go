package usecase

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/entity"
)

type UserUseCase struct {
	UserInterface entity.UserInterface
}

func NewUser(
	UserInterface entity.UserInterface,
) *UserUseCase {
	return &UserUseCase{
		UserInterface: UserInterface,
	}
}

func (a *UserUseCase) CreateNewUser(input dto.UserDTO) (dto.UserDTO, error) {
	User := entity.UserEntity{
		ID:         input.ID,
		Name:       input.Name,
		Password:   input.Password,
		Email:      input.Email,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		AcountType: input.AcountType,
	}
	if _, err := a.UserInterface.CreateUser(&User); err != nil {
		return dto.UserDTO{}, err
	}
	dto := dto.UserDTO{
		ID:         User.ID,
		Name:       input.Name,
		Password:   input.Password,
		Email:      input.Email,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		AcountType: input.AcountType,
	}

	return dto, nil
}

func (a *UserUseCase) FindAllUsers(page, limit int, sort string) ([]dto.UserDTO, error) {
	entities, err := a.UserInterface.FindAllUsers(page, limit, sort)
	if err != nil {
		return nil, err
	}
	var User []dto.UserDTO
	for _, entity := range entities {
		dto := dto.UserDTO{
			ID:         entity.ID,
			Name:       entity.Name,
			Password:   entity.Password,
			Email:      entity.Email,
			Birthday:   entity.Birthday,
			Gender:     entity.Gender,
			AcountType: entity.AcountType,
		}
		User = append(User, dto)
	}

	return User, nil
}

func (a *UserUseCase) FindUserByID(id string) (dto.UserDTO, error) {
	entity, err := a.UserInterface.FindUserByID(id)
	if err != nil {
		return dto.UserDTO{}, err
	}
	dto := dto.UserDTO{
		ID:         entity.ID,
		Name:       entity.Name,
		Password:   entity.Password,
		Email:      entity.Email,
		Birthday:   entity.Birthday,
		Gender:     entity.Gender,
		AcountType: entity.AcountType,
	}
	return dto, nil
}
