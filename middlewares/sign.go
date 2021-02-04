package middlewares

import (
	"net/http"

	"github.com/dasha-kinsely/Countryside/entities/apis"
	"github.com/dasha-kinsely/Countryside/helpers"
	"github.com/gin-gonic/gin"
)

func Sign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := apis.Result{}
		sign, err := helpers.VerifyTimeSignBasic(ctx)
		if sign != nil {
			res.SetCode(1)
			res.SetMessage("")
			res.SetData(sign)
			ctx.JSON(http.StatusUnauthorized, res)
			ctx.Abort()
			return
		}
		if err != nil {
			res.SetCode(1)
			res.SetMessage(err.Error())
			ctx.JSON(http.StatusUnauthorized, res)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
