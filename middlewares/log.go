package middlewares

import (
	"github.com/dasha-kinsely/Countryside/configs"
	"github.com/gin-gonic/gin"
)

func LogToFile() gin.HandlerFunc {
	logFilePath := configs.LOG_FILE_PATH
	logFileName := configs.LOG_FILE_NAME
}
