package logs

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// setupLogOutput opens a new log file with os std multiwriter
func setup(fileType string) {
	if fileType == "e" {
		f, _ := os.Create("error.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}
