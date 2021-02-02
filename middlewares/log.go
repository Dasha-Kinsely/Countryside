package middlewares

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Dasha-Kinsely/Countryside/configs"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// LogToFile is a collection of functions that handle file-based logging
func LogToFile() gin.HandlerFunc {
	logFilePath := configs.LOG_PATH
	logFileName := configs.LOG_FILE_NAME
	fileName := path.Join(logFilePath, logFileName)
	// Open Logger Files
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// Third Party Logrus Standardized Logging
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	// Third Party File-Rotatelogs to Split Log Files
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	// Third Party Lfshook to Standardize Various Logs Using JSON Formatter
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logger.AddHook(lfHook)
	// The final product is a bundle of log output below
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqURI := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURI,
		}).Info()
	}
}
