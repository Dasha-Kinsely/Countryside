package entities

//gorm:"primary_key;auto_increment"
// binding:"required,url" gorm:"type:varchar(256);UNIQUE"
type Farm struct {
	ID          uint   `json:"id" binding:"required"`
	URL         string `json:"url" binding:"required,url"`
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Description string `json:"description"`
	Admin       User   `json:"admin"`
}
