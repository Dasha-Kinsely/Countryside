package models

type UserModel struct {
	BaseModel
	ID       uint   `gorm:"primary_key"`
	Email    string `gorm:"column:email;unique"`
	Username string `gorm:"column:username;unique"`
}
