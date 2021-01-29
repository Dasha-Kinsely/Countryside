package logs

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// setupLogOutput opens a new log file with os std multiwriter
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
