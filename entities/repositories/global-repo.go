package repositories

import (
	"github.com/jinzhu/gorm"
)

type database struct {
	connection *gorm.DB
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database...\n")
	}
}
