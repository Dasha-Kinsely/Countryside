package entities

//gorm:"primary_key;auto_increment"
// binding:"required,url" gorm:"type:varchar(256);UNIQUE"
type Farm struct {
	ID          uint   `json:"id" `
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Admin       string `json:"admin"`
}
