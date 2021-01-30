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
	server := gin.New()
	server.Use(gin.Recovery(), logs.HTTPLogger())
	db, initDBErr := configs.initDB()
	if initDBErr != nil {

	}
	server.GET("/test/init", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
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
	/*server.POST("/binder", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, )
	})*/

	server.Run(":8080")
}
