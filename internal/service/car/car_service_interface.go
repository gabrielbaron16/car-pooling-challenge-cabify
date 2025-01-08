package car

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type ICarService interface {
	ResetCars(cars []*model.Car) error
}

var instance ICarService

func GetInstance() ICarService {
	return instance
}

func SetInstance(service ICarService) {
	instance = service
}
