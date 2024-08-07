package usecase_test

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func generatePersonalTrainer(f *fuzz.Fuzzer) *dtos.User {
	personalTrainer := &dtos.User{}
	f.Fuzz(personalTrainer)
	personalTrainer.ID = uuid.New().String()
	personalTrainer.ApplicationId = uuid.New().String()
	f.Fuzz(&personalTrainer.Name)
	f.Fuzz(&personalTrainer.Charge)
	f.Fuzz(&personalTrainer.Gym)
	return personalTrainer
}
func TestCreatePersonalTrainer(t *testing.T){
	f := fuzz.New()
	personal := generatePersonalTrainer(f)

	createPersonalResponse, err := GymUseCase.CreatePersonal(personal)
	assert.Nil(t, err)
	assert.Equal(t, personal.Name, createPersonalResponse.ID)
	assert.Equal(t, personal.Name, createPersonalResponse.Name)
	assert.Equal(t, entity.UserEnum.Personal, createPersonalResponse.Charge)
	assert.Equal(t, personal.Name, createPersonalResponse.Gym)
}