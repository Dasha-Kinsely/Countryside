package helpers

import (
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeSignBasic(ctx *gin.Context) {
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	responseMapping := map[string]interface{}{}
	params := url.Values{
		"title": []string{"received..."},
		"time":  []string{timeStamp},
	}
	responseMapping["time_sign_basic"] = CreateTimeSignBasic(params)
	ReturnedJSON("200", "", responseMapping, ctx)
}
