package configs

import (
	"Countryside/configs/confEntities"
	"fmt"

	"github.com/go-ini/ini"
)

func LoadEnv() {
	var cfg *ini.File
	var err error
	cfg, err = ini.Load("ini/app.ini")
	if err != nil {
		fmt.Println("Error has occurred upon loading ENV file...\n")
	}
	mapIniFileToStruct(cfg, "database", &confEntities.Database{})
	mapIniFileToStruct(cfg, "server", &confEntities.Server{})
}

func mapIniFileToStruct(cfg *ini.File, iniLine string, structVar interface{}) {
	err := cfg.Section(iniLine).MapTo(structVar)
	if err != nil {
		fmt.Println("Error has occurred while mapping ENV...\n")
	}
}
