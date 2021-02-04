package main

import (
	"github.com/Dasha-Kinsely/Countryside/configs"
	"github.com/Dasha-Kinsely/Countryside/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	routers.InitRouter(engine)
	engine.Run(configs.PORT)
}
