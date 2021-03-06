package entities

type User struct {
	Base
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"required" binding:"email" gorm:"primary_key"`
}
