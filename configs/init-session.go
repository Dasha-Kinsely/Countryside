package configs

import (
	"github.com/gin-gonic/gin"
)

func SetSession(session *gin.Engine) string {
	return "done configuring session...\n"
}
