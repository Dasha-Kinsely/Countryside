package entities

type User struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" validate:"required" binding:"email"`
}
