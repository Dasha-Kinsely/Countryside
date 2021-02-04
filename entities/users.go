package entities

type User struct {
	Name string `form:"name" json:"name" binding:"required,NameValid"`
}
