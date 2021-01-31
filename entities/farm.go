package entities

type Farm struct {
	Base
	ID          uint64 `json:"id" binding:"required" gorm:"primary_key;auto_increment"`
	URL         string `json:"url" binding:"required,url" gorm:"type:varchar(255);UNIQUE"`
	Name        string `json:"name" binding:"required,min=2,max=50" gorm:"type:varchar(155)"`
	Description string `json:"description" gorm:"type:text"`
	Admin       User   `json:"admin" gorm:"foreignkey:UserEmail"`
	UserEmail   string `json:"-"`
}
