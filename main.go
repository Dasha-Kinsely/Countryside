package main

import (
	"Countryside/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadEnv()
	db := configs.InitDB()
	configs.MigrateDB()
	defer db.Close()
	server := gin.Default()
	server.Run(":8080")
}
