package controllers

import (
	"strconv"

	"github.com/Dasha-Kinsely/Countryside/entities"
	"github.com/Dasha-Kinsely/Countryside/services"
	"github.com/gin-gonic/gin"
)

type FarmController interface {
	Save(ctx *gin.Context) error
	FindAll() []entities.Farm
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

func (c *controller) Update(ctx *gin.Context) error {
	var farm entities.Farm
	err := ctx.ShouldBindJSON(&farm)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	farm.ID = id
	c.service.Update(farm)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var farm entities.Farm
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	farm.ID = id
	c.service.Delete(farm)
	return nil
}
