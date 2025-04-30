package usecase

import (
	"errors"

	"github.com/mathgod152/GymControl/internal/dto"
	"github.com/mathgod152/GymControl/internal/entity"
)

type PersonalUsecase struct {
	PersonalUsecase entity.PersonalInteractorInterface
}

func (p *PersonalUsecase) AssignCustomerToTrainer(userEmail, personalEmail string) error {
	isPersonal, err := p.PersonalUsecase.IsTrainer(personalEmail)
	if err != nil || !isPersonal {
		return errors.New("user is not a personal")
	}
	if err := p.PersonalUsecase.AssociateUserWithTrainer(userEmail, personalEmail); err != nil {
		return errors.New("Erro to Relationate User")
	}
	return nil
}

func (p *PersonalUsecase) FindTrainerCustomers(personalEmail string) ([]dto.CustumerDto, error) {
	isPersonal, err := p.PersonalUsecase.IsTrainer(personalEmail)
	if err != nil || !isPersonal {
		return nil, errors.New("user is not a personal")
	}
	customersReturn, err := p.PersonalUsecase.FindUsersByTrainer(personalEmail)
	if err != nil {
		return nil, errors.New("error to find custumers")
	}
	var customers []dto.CustumerDto
	for _, entity := range customersReturn {
		customers = append(customers, dto.CustumerDto{
			Name:       entity.Name,
			Email:      entity.Email,
			Gender:     entity.Gender,
		})
	}
	return customers, nil
}

func (p *PersonalUsecase) FindCustomerTrainer(personalEmail string) ([]dto.PersonalDto, error) {
	customersReturn, err := p.PersonalUsecase.FindTrainersByUser(personalEmail)
	if err != nil {
		return nil, errors.New("error to find personal")
	}
	var customers []dto.PersonalDto
	for _, entity := range customersReturn {
		customers = append(customers, dto.PersonalDto{
			Name:       entity.Name,
			Email:      entity.Email,
			Birthday:   entity.Birthday,
			Gender:     entity.Gender,
		})
	}
	return customers, nil
}
