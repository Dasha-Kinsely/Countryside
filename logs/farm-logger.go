package logs

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTPLogger Basic Info
func HTTPLogger() gin.HandlerFunc {
	setup("p")
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"CLIENT: %s - [%s]\n PATH: %s -> %s -> %s -> %d\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Path,
			param.Method,
			param.Latency,
			param.StatusCode)
	})
}

func ErrorLogger() gin.HandlerFunc {
	setup("e")
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"CLIENT: %s - [%s]\n PATH: %s -> %s -> %s -> %d\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Path,
			param.Method,
			param.Latency,
			param.StatusCode)
	})
}
