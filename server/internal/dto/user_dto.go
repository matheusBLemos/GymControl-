package dto

type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	AcountType string `json:"acountType"`
}
