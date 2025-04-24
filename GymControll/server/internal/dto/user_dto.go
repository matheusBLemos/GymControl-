package dto

type CreateUserDto struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	Gender     string `json:"gender"`
	AcountType string `json:"acountype"`
}

type UserDto struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	Gender     string `json:"gender"`
	AcountType string `json:"acountype"`
}