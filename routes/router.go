package routes

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

// RunRouter is the entry point
func RunRouter(r *gin.Engine) {
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
}
