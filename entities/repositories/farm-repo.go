package repositories

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type FarmRepository interface {
	SaveFarm(farm entities.Farm)
	UpdateFarm(farm entities.Farm)
	DeleteFarm(farm entities.Farm)
	FinalAllFarms() []entities.Farm
	CloseDBFarm()
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
