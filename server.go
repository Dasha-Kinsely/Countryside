package main

import (
	"net/http"

	"github.com/Dasha-Kinsely/Countryside/configs"
	"github.com/Dasha-Kinsely/Countryside/controllers"
	"github.com/Dasha-Kinsely/Countryside/logs"
	"github.com/Dasha-Kinsely/Countryside/services"
	"github.com/gin-gonic/gin"
)

var (
	farmService    services.FarmService       = services.New()
	farmController controllers.FarmController = controllers.New(farmService)
)

func main() {
	// Init Gin Server
	server := gin.New()
	gin.SetMode(gin.DebugMode)
	// Init Logger Middlewares
	server.Use(gin.Recovery(), logs.HTTPLogger())
	// Init DB using Primitive Methods
	db, initDBErr := configs.InitDB()
	if initDBErr != nil {
		logs.ErrorLogger("db failed")
		return
	}
	defer db.Close()
	// Init Router with Configs
	router := gin.Default()
	configs.SetTemplate(router)
	configs.SetSession(router)

	server.GET("/farms", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, farmController.FindAll())
	})
	server.POST("/farms", func(ctx *gin.Context) {
		err := farmController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "valid input"})
		}
	})

	server.Run(":8080")
}
