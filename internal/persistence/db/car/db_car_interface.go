package car

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

var carDAO CarDatabase

type CarDatabase interface {
	GetCarById(carId uint) *model.Car
	UpsertCar(car *model.Car)
	GetAllCars() map[uint]*model.Car
	ResetCars()
}

func GetInstance() CarDatabase {
	return carDAO
}

func SetInstance(instance CarDatabase) {
	carDAO = instance
}
