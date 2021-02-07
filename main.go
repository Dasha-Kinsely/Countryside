package main

import (
	"Countryside/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.Init()
	configs.MigrateDB()
	defer db.Close()
	server := gin.Default()
	server.Run(":8080")
}
