package main

import (
	"github.com/Dasha-Kinsely/Countryside/configs"
	"github.com/Dasha-Kinsely/Countryside/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	routers.InitRouter(engine)
	engine.Run(configs.PORT)
}
