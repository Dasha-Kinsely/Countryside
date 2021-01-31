package routes

import (
	"net/http"

	"github.com/Dasha-Kinsely/Countryside/controllers"
	"github.com/Dasha-Kinsely/Countryside/entities/repositories"
	"github.com/Dasha-Kinsely/Countryside/services"
	"github.com/gin-gonic/gin"
)

var (
	farmRepository repositories.FarmRepository = repositories.NewFarmRepository()
	farmService    services.FarmService        = services.New(farmRepository)
	farmController controllers.FarmController  = controllers.New(farmService)
)

// RunRouter is the entry point
func RunRouter(r *gin.Engine) {
	defer farmRepository.CloseDBFarm()

	r.GET("/farms", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, farmController.FindAll())
	})
	r.POST("/farms", func(ctx *gin.Context) {
		err := farmController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "valid input"})
		}
	})
	r.PUT("/farms/:id", func(ctx *gin.Context) {
		err := farmController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "valid update"})
		}
	})
	r.DELETE("/farms/:id", func(ctx *gin.Context) {
		err := farmController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "valid delete"})
		}
	})
}
