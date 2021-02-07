package configs

import (
	"Countryside/configs/confEntities"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Pointer for DB
var DB *gorm.DB

// Initialize DB
func InitDB() *gorm.DB {
	dbConfig := &confEntities.Database{}
	dbRawStruct := []string{dbConfig.User, ":", dbConfig.Password, "@tcp(", dbConfig.Host, ")/", dbConfig.DatabaseName, "?charset=utf8&parseTime=True&loc=Local"}
	dbString := configStringBuilder(dbRawStruct)
	fmt.Println(dbString)
	db, err := gorm.Open(dbConfig.Type, dbString)
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

func configStringBuilder(array []string) string {
	var b strings.Builder
	for _, v := range array {
		_, err := b.WriteString(v)
		if err != nil {
			fmt.Println("String Builder Error: ", err)
		}
	}
	return b.String()
}
