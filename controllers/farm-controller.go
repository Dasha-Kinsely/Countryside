package controllers

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
	"github.com/Dasha-Kinsely/Countryside/services"
	"github.com/gin-gonic/gin"
)

type FarmController interface {
	Save(ctx *gin.Context) error
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
func (c *controller) Save(ctx *gin.Context) error {
	var farm entities.Farm
	err := ctx.ShouldBindJSON(&farm)
	if err != nil {
		return err
	}
	c.service.Save(farm)
	return nil
}
