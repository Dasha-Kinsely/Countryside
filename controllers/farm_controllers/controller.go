package farm_controllers

import (
	"github.com/gin-gonic/gin"
)

// AddFarm w/ default admin
func AddFarm(ctx *gin.Context) {
	name := ctx.Query("name")
	admin := ctx.DefaultQuery("admin", "none")
	ctx.JSON(200, gin.H{
		"farm":  "AddFarm",
		"name":  name,
		"admin": admin,
	})
}
