package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Dasha-Kinsely/Countryside/configs"
	"github.com/gin-gonic/gin"
)

func VerifyTimeSignBasic(ctx *gin.Context) {
	fmt.Println("Verifying")
	var method = ctx.Request.Method
	var timestamp int64
	//var req url.Values
	// Make sure method is not illegal, then get the corresponding data
	if method == "GET" {
		//req = ctx.Request.URL.Query()
		//sign = ctx.Query("sn")
		timestamp, _ = strconv.ParseInt(ctx.Query("ts"), 10, 64)
	} else if method == "POST" {
		ctx.Request.ParseForm()
		//req = ctx.Request.PostForm
		//sign = ctx.PostForm("sn")
		timestamp, _ = strconv.ParseInt(ctx.PostForm("ts"), 10, 64)
	} else {
		ReturnedJSON("500", "illegal method", "", ctx)
		return
	}
	// Make sure it hasn't expired
	exp, _ := strconv.ParseInt(configs.API_EXPIRY, 10, 64)
	if timestamp > time.Now().Unix() || time.Now().Unix()-timestamp >= exp {
		ReturnedJSON("500", "Timestamp error", "", ctx)
		return
	}
	// Make sure time stamp is not none or different from basic signs
	/*if sign == "" || sign != CreateTimeSignBasic(req) {
		ReturnedJSON("500", "Sign srror", "", ctx)
		return
	}*/
}
