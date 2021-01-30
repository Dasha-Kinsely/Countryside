package configs

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB Pointer from gorm package
var DB *gorm.DB

// InitDB using primitive methods
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "yuntai:dkemployeeaccess@tcp(localhost:3306)/masterdb?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		db.AutoMigrate(&entities.Farm{}, &entities.User{})
		return db, err
	}
	return nil, err
}
