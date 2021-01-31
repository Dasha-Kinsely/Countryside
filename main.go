package main

import (
	"github.com/Dasha-Kinsely/Countryside/logs"
	"github.com/Dasha-Kinsely/Countryside/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init Gin Server
	server := gin.New()
	// Init Logger Middlewares
	server.Use(gin.Recovery(), logs.HTTPLogger())
	// Init DB using Primitive Methods
	/*db, initDBErr := configs.InitDB()
	if initDBErr != nil {
		logs.ErrorLogger("db failed")
		return
	}
	defer db.Close()*/
	// Main Operations on the Server & Router & Wrapper
	routes.RunRouter(server)
	server.Run("localhost:8082")
}
