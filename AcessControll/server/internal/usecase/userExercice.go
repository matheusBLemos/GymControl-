package usecase

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/entity"
)

type UserTrainingUseCase struct {
	UserTrainingInterface entity.UserTrainingInterface
}

func NewUserTraining(
	UserTrainingInterface entity.UserTrainingInterface,
) *UserTrainingUseCase {
	return &UserTrainingUseCase{
		UserTrainingInterface: UserTrainingInterface,
	}
}

func (a *UserTrainingUseCase) CreateNewUserTraining(input dto.UserTrainingDto) (dto.UserTrainingDto, error) {
	UserTraining := entity.UserTrainingEntity{
		ID:           input.ID,
		IdUser:       input.IdUser,
		IdTraining:   input.IdTraining,
		IdCreateUser: input.IdCreateUser,
	}
	if _, err := a.UserTrainingInterface.Create(&UserTraining); err != nil {
		return dto.UserTrainingDto{}, err
	}
	dto := dto.UserTrainingDto{
		ID:           UserTraining.ID,
		IdUser:       input.IdUser,
		IdTraining:   input.IdTraining,
		IdCreateUser: input.IdCreateUser,
	}

	return dto, nil
}

func (a *UserTrainingUseCase) FindAllUserTrainings(page, limit int, sort string) ([]dto.UserTrainingDto, error) {
	entities, err := a.UserTrainingInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}
	var User []dto.UserTrainingDto
	for _, entity := range entities {
		dto := dto.UserTrainingDto{
			ID:           entity.ID,
			IdUser:       entity.IdUser,
			IdTraining:   entity.IdTraining,
			IdCreateUser: entity.IdCreateUser,
		}
		User = append(User, dto)
	}

	return User, nil
}

func (a *UserTrainingUseCase) FindUserTrainingByID(id string) (dto.UserTrainingDto, error) {
	entity, err := a.UserTrainingInterface.FindByID(id)
	if err != nil {
		return dto.UserTrainingDto{}, err
	}
	dto := dto.UserTrainingDto{
		ID:           entity.ID,
		IdUser:       entity.IdUser,
		IdCreateUser: entity.IdCreateUser,
		IdTraining:   entity.IdTraining,
	}
	return dto, nil
}
