package repositories

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserRepository interface {
	SaveUser(user entities.User)
	UpdateUser(user entities.User)
	DeleteUser(user entities.User)
	FinalAllUsers() []entities.User
}

func (db *database) SaveUser(user entities.User) {
	db.connection.Create(&user)
}

func (db *database) UpdateUser(user entities.User) {
	db.connection.Update(&user)
}

func (db *database) DeleteUser(user entities.User) {
	db.connection.Delete(&user)
}

func (db *database) FindAllUsers() []entities.User {
	var user []entities.User
	db.connection.Set("gorm:auto_preload", true).Find(&user)
	return user
}
