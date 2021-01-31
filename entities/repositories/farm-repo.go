package repositories

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type FarmRepository interface {
	SaveFarm(farm entities.Farm)
	UpdateFarm(farm entities.Farm)
	DeleteFarm(farm entities.Farm)
	FindAllFarms() []entities.Farm
	CloseDBFarm()
}
type database struct {
	connection *gorm.DB
}

func NewFarmRepository() FarmRepository {
	db, err := gorm.Open("mysql", "yuntai:dkemployeeaccess@tcp(localhost:3306)/masterdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to open")
	}
	db.AutoMigrate(&entities.Farm{}, &entities.User{})
	return &database{
		connection: db,
	}
}

func (db *database) SaveFarm(farm entities.Farm) {
	db.connection.Create(&farm)
}

func (db *database) UpdateFarm(farm entities.Farm) {
	db.connection.Update(&farm)
}

func (db *database) DeleteFarm(farm entities.Farm) {
	db.connection.Delete(&farm)
}

func (db *database) FindAllFarms() []entities.Farm {
	var farm []entities.Farm
	db.connection.Set("gorm:auto_preload", true).Find(&farm)
	return farm
}

func (db *database) CloseDBFarm() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close db")
	}
}
