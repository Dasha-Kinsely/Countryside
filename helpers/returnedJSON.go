package helpers

import (
	"github.com/gin-gonic/gin"
)

func ReturnedJSON(code, msg string, data interface{}, ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	ctx.Abort()
}
