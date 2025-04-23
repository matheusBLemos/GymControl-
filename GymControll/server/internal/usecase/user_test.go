package usecase_test

// import (
// 	"testing"
// 	fuzz "github.com/google/gofuzz"
// )

// func CreateNewUserTest(t *testing.T) {
// 	f := fuzz.New()

// 	newUser := &dto.User{
// 		f.Fuzz(&newUser.Id)
// 		f.Fuzz(&newUser.Name)
// 		f.Fuzz(&newUser.Password)
// 		f.Fuzz(&newUser.Email)
// 		f.Fuzz(&newUser.Birthday)
// 		f.Fuzz(&newUser.Gender)
// 		f.Fuzz(&newUser.AcountType)		
// 	}
	
// 	createUserResponse, err := usecase.UserUsecase.NewUser(NewUser)
// 	assert.Equal(t, createUserResponse.Id, createUserResponse.Id)
// 	assert.Equal(t, createUserResponse.Name, createUserResponse.Name)
// 	assert.Equal(t, createUserResponse.Password, createUserResponse.Password)
// 	assert.Equal(t, createUserResponse.Email, createUserResponse.Email)
// 	assert.Equal(t, createUserResponse.Birthday, createUserResponse.Birthday)
// 	assert.Equal(t, createUserResponse.Gender, createUserResponse.Gender)
// 	assert.Equal(t, createUserResponse.AcountType, createUserResponse.AcountType)
// 	assert.NoError(t, err)
// }


