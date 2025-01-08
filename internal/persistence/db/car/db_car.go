package car

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type CarDbImp struct {
	Cars map[uint]*model.Car
}

func (c CarDbImp) GetCarById(carID uint) *model.Car {
	return c.Cars[carID]
}

func (c CarDbImp) UpsertCar(car *model.Car) {
	c.Cars[car.ID] = car
}

func (c CarDbImp) GetAllCars() map[uint]*model.Car {
	return c.Cars
}

func (c CarDbImp) ResetCars() {
	c.Cars = make(map[uint]*model.Car)
}
