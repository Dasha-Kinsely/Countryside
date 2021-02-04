package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Dasha-Kinsely/Countryside/configs"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func VerifyTimeSignBasic(ctx *gin.Context) (res map[string]string, err error){
	fmt.Println("Verifying")
	var method = ctx.Request.Method
	var timestamp int64
	var debug string
	//var req url.Values
	// Make sure method is not illegal, then get the corresponding data
	if method == "GET" {
		//req = ctx.Request.URL.Query()
		//sign = ctx.Query("sn")
		debug = ctx.Query("debug")
		timestamp, _ = strconv.ParseInt(ctx.Query("ts"), 10, 64)
	} else if method == "POST" {
		ctx.Request.ParseForm()
		//req = ctx.Request.PostForm
		//sign = ctx.PostForm("sn")
		debug = ctx.PostForm("debug")
		timestamp, _ = strconv.ParseInt(ctx.PostForm("ts"), 10, 64)
	} else {
		//ReturnedJSON("500", "illegal method", "", ctx)
		err = errors.New("Illegal Methods")
		return
	}
	ts := time.Now().Unix()
	if debug == "1" {
		res = map[string]string{
			"ts": strconv.FormatInt(ts, 10)
		}
		return
	}
	// Make sure it hasn't expired
	expiry, _ := strconv.ParseInt(configs.API_EXPIRY, 10, 64)
	if timestamp > ts || ts-timestamp >= expiry {
		err = errors.New("--ERROR: Timestamp error--")
		return
	}
	return
	// Make sure time stamp is not none or different from basic signs
	/*if sign == "" || sign != CreateTimeSignBasic(req) {
		ReturnedJSON("500", "Sign srror", "", ctx)
		return
	}*/
}
