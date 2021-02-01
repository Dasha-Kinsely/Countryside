package user_controllers

import (
	"github.com/gin-gonic/gin"
)

// AddUser adds a user
func AddUser(ctx *gin.Context) {
	name := ctx.Query("name")
	ctx.JSON(200, gin.H{
		"user": "AddUser",
		"name": name,
	})
}
