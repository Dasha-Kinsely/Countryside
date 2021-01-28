package controllers

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
	"github.com/Dasha-Kinsely/Countryside/services"
	"github.com/gin-gonic/gin"
)

type FarmController interface {
	Save(ctx *gin.Context) entities.Farm
	FindAll() []entities.Farm
}
type controller struct {
	service services.FarmService
}

func New(service services.FarmService) FarmController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entities.Farm {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entities.Farm {
	var farm entities.Farm
	ctx.BindJSON(&farm)
	c.service.Save(farm)
	return farm
}
