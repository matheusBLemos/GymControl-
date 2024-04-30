package usecase

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/entity"
)

type TrainingUseCase struct {
	TrainingInterface entity.TrainingInterface
}

func NewTraining(
	TrainingInterface entity.TrainingInterface,
) *TrainingUseCase {
	return &TrainingUseCase{
		TrainingInterface: TrainingInterface,
	}
}

func (a *TrainingUseCase) CreateNewTraining(input dto.TrainingDto) (dto.TrainingDto, error) {
	Training := entity.TrainingEntity{
		ID:           input.ID,
		IdUser:       input.IdUser,
		IdCreateUser: input.IdCreateUser,
		IdTraining:   input.IdTraining,
		Desc:         input.Desc,
	}
	if _, err := a.TrainingInterface.Create(&Training); err != nil {
		return dto.TrainingDto{}, err
	}
	dto := dto.TrainingDto{
		ID:           Training.ID,
		IdUser:       input.IdUser,
		IdCreateUser: input.IdCreateUser,
		IdTraining:   input.IdTraining,
		Desc:         input.Desc,
	}

	return dto, nil
}

func (a *TrainingUseCase) FindAllTrainings(page, limit int, sort string) ([]dto.TrainingDto, error) {
	entities, err := a.TrainingInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}
	var User []dto.TrainingDto
	for _, entity := range entities {
		dto := dto.TrainingDto{
			ID:           entity.ID,
			IdUser:       entity.IdUser,
			IdCreateUser: entity.IdCreateUser,
			IdTraining:   entity.IdTraining,
			Desc:         entity.Desc,
		}
		User = append(User, dto)
	}

	return User, nil
}

func (a *TrainingUseCase) FindTrainingByID(id string) (dto.TrainingDto, error) {
	entity, err := a.TrainingInterface.FindByID(id)
	if err != nil {
		return dto.TrainingDto{}, err
	}
	dto := dto.TrainingDto{
		ID:           entity.ID,
		IdUser:       entity.IdUser,
		IdCreateUser: entity.IdCreateUser,
		IdTraining:   entity.IdTraining,
		Desc:         entity.Desc,
	}
	return dto, nil
}
