package usecase

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/entity"
)

type ExerciseUseCase struct {
	ExerciseInterface entity.ExerciceInterface
}

func NewExercise(
	ExerciseInterface entity.ExerciceInterface,
) *ExerciseUseCase {
	return &ExerciseUseCase{
		ExerciseInterface: ExerciseInterface,
	}
}

func (a *ExerciseUseCase) CreateNewExercise(input dto.ExerciceDto) (dto.ExerciceDto, error) {
	Exercise := entity.ExerciceEntity{
		ID:           input.ID,
		Name:         input.Name,
		Desc:         input.Desc,
		CreateUser:   input.CreateUser,
		ExerciceName: input.ExerciceName,
	}
	if _, err := a.ExerciseInterface.Create(&Exercise); err != nil {
		return dto.ExerciceDto{}, err
	}
	dto := dto.ExerciceDto{
		ID:           Exercise.ID,
		Name:         input.Name,
		Desc:         input.Desc,
		CreateUser:   input.CreateUser,
		ExerciceName: input.ExerciceName,
	}

	return dto, nil
}

func (a *ExerciseUseCase) FindAllExercises(page, limit int, sort string) ([]dto.ExerciceDto, error) {
	entities, err := a.ExerciseInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}
	var User []dto.ExerciceDto
	for _, entity := range entities {
		dto := dto.ExerciceDto{
			ID:           entity.ID,
			Name:         entity.Name,
			Desc:         entity.Desc,
			CreateUser:   entity.CreateUser,
			ExerciceName: entity.Desc,
		}
		User = append(User, dto)
	}

	return User, nil
}

func (a *ExerciseUseCase) FindExerciseByID(id string) (dto.ExerciceDto, error) {
	entity, err := a.ExerciseInterface.FindByID(id)
	if err != nil {
		return dto.ExerciceDto{}, err
	}
	dto := dto.ExerciceDto{
		ID:           entity.ID,
		Name:         entity.Name,
		Desc:         entity.Desc,
		CreateUser:   entity.CreateUser,
		ExerciceName: entity.Desc,
	}
	return dto, nil
}
