package entities

type Farm struct {
	Name  string `form:"name" json:"name" binding:"required,NameValid"`
	Admin string `form:"admin" json:"admin" binding:"required"`
}
