package services

import (
	"github.com/Dasha-Kinsely/Countryside/entities"
	"github.com/Dasha-Kinsely/Countryside/entities/repositories"
)

/*FarmService Interface*/
type FarmService interface {
	Save(entities.Farm) entities.Farm
	FindAll() []entities.Farm
	Update(entities.Farm)
	Delete(entities.Farm)
}

/*Actual Body of entities.Farm*/
type farmService struct {
	farmRepo repositories.FarmRepository
}

/*New farmService struct init*/
func New(repo repositories.FarmRepository) FarmService {
	return &farmService{
		farmRepo: repo,
	}
}

func (service *farmService) Save(farm entities.Farm) entities.Farm {
	service.farmRepo.SaveFarm(farm)
	return farm
}

func (service *farmService) FindAll() []entities.Farm {
	return service.farmRepo.FindAllFarms()
}

func (service *farmService) Update(farm entities.Farm) {
	service.farmRepo.UpdateFarm(farm)
}

func (service *farmService) Delete(farm entities.Farm) {
	service.farmRepo.DeleteFarm(farm)
}
