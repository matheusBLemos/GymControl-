package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mathgod152/GymControl/internal/dto"
	"github.com/mathgod152/GymControl/internal/entity"
	"github.com/mathgod152/GymControl/internal/usecase"
)

// Mock para UserInterface
type mockUserInterface struct {
	PasswordHash string 
	Users     []*entity.User
	CreateErr error
}

func newMockUserInterface() *mockUserInterface {
	return &mockUserInterface{
		Users: []*entity.User{
			{Id: "id1", Name: "João", Email: "joao@example.com"},
			{Id: "id2", Name: "Maria", Email: "maria@example.com"},
			{Id: "id3", Name: "Carlos", Email: "carlos@example.com"},
		},
	}
}


func (m *mockUserInterface) CryptoPassword(password string) (string, error) {
	return "hashedpassword", nil
}

func (m *mockUserInterface) Create(user *entity.User) (*entity.User, error) {
	return user, nil
}
func (m *mockUserInterface) FindAll(page, limit int, sort string) ([]*entity.User, error) {
	return m.Users, nil
}

func (m *mockUserInterface) FindByID(id string) (*entity.User, error) {
	for _, user := range m.Users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *mockUserInterface) Update(user *entity.User) (*entity.User, error) {
	for i, u := range m.Users {
		if u.Id == user.Id {
			m.Users[i] = user
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *mockUserInterface) Delete(id string) (bool, error) {
	for i, u := range m.Users {
		if u.Id == id {
			m.Users = append(m.Users[:i], m.Users[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("user not found")
}

// Mock para UserPersonalInterface (não usado nesse caso, mas necessário pra construir o usecase)
type mockUserPersonalInterface struct {
	Relations map[string][]*entity.User // mapa de personalID para usuários
	Trainer   map[string][]*entity.User // mapa de userID para trainers
}

func newMockUserPersonalInterface() *mockUserPersonalInterface {
	return &mockUserPersonalInterface{
		Relations: map[string][]*entity.User{
			"personal1": {
				{Id: "user1", Name: "Aluno 1"},
				{Id: "user2", Name: "Aluno 2"},
			},
			"personal2": {
				{Id: "user3", Name: "Aluno 3"},
			},
		},
		Trainer: map[string][]*entity.User{
			"user1": {
				{Id: "personal1", Name: "Personal João"},
			},
			"user2": {
				{Id: "personal1", Name: "Personal João"},
			},
			"user3": {
				{Id: "personal2", Name: "Personal Maria"},
			},
		},
	}
}

func (m *mockUserPersonalInterface) Relationate(userEmail, personalEmail string) (*entity.User, error) {
	// Simula uma relação criada com sucesso
	return &entity.User{
		Email: userEmail,
	}, nil
}

func (m *mockUserPersonalInterface) FindUserByTrainer(personalId string) ([]*entity.User, error) {
	if users, ok := m.Relations[personalId]; ok {
		return users, nil
	}
	return []*entity.User{}, nil
}

func (m *mockUserPersonalInterface) FindTrainerByUser(userId string) ([]*entity.User, error) {
	if trainers, ok := m.Trainer[userId]; ok {
		return trainers, nil
	}
	return []*entity.User{}, nil
}

func FuzzCreateNewUser(f *testing.F) {
	// Semente inicial
	f.Add("user@example.com", "123456", "João", "male", "basic")

	f.Fuzz(func(t *testing.T,
		email string,
		password string,
		name string,
		gender string,
		acountType string,
	) {
		u := usecase.UserUsecase{
			UserInterface: &mockUserInterface{
				PasswordHash: "hashedPassword",
			},
			UserPersonalInterface: &mockUserPersonalInterface{},
		}

		dtoInput := dto.CreateUserDto{
			Name:       name,
			Email:      email,
			Password:   password,
			Birthday:   time.Now().Format("2006-01-02"),
			Gender:     gender,
			AcountType: acountType,
		}
 
		_, err := u.CreateNewUser(dtoInput)

		// Você pode checar que não houve panic e que os erros esperados são retornados apropriadamente
		if err != nil &&
			err.Error() != "Error to Encrypt Password" &&
			err.Error() != "Error to create User" {
			t.Logf("Possível erro de validação: %v", err)
		}
	})
}
