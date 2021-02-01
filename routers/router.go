package routers

import (
	"github.com/Dasha-Kinsely/Countryside/controllers/farm_controllers"
	"github.com/Dasha-Kinsely/Countryside/controllers/user_controllers"
	"github.com/Dasha-Kinsely/Countryside/helpers"
	"github.com/gin-gonic/gin"
)

// Entry Point of Routing
func InitRouter(r *gin.Engine) {
	r.GET("/sign", helpers.TimeSignBasic)
	FarmGroup := r.Group("/farm", helpers.VerifyTimeSignBasic)
	{
		FarmGroup.Any("/add", farm_controllers.AddFarm)
	}
	UserGroup := r.Group("/user")
	{
		UserGroup.Any("/add", user_controllers.AddUser)
	}
}
