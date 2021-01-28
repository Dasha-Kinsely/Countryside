package services

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
)

/*FarmService Interface*/
type FarmService interface {
	Save(entities.Farm) entities.Farm
	FindAll() []entities.Farm
}

/*Actual Body of entities.Farm*/
type farmService struct {
	farms []entities.Farm
}

/*New farmService struct init*/
func New() FarmService {
	return &farmService{}
}

func (service *farmService) Save(farm entities.Farm) entities.Farm {
	service.farms = append(service.farms, farm)
	return farm
}

func (service *farmService) FindAll() []entities.Farm {
	return service.farms
}
