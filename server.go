package main

import (
	"net/http"

	"github.com/Dasha-Kinsely/Countryside/controllers"
	"github.com/Dasha-Kinsely/Countryside/services"
	"github.com/gin-gonic/gin"
)

var (
	farmService    services.FarmService       = services.New()
	farmController controllers.FarmController = controllers.New(farmService)
)

func main() {
	server := gin.Default()
	server.GET("/test/init", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	server.GET("/farms", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, farmController.FindAll())
	})
	server.POST("/farms", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, farmController.Save(ctx))
	})

	server.Run(":8080")
}
