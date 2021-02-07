package configs

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Pointer for DB
var DB *gorm.DB

// Initialize DB
func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "yuntai:dkemployeeaccess@tcp(localhost:3306)/masterdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db error: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

// Return Active DB by Its Address
func GetDB() *gorm.DB {
	return DB
}
