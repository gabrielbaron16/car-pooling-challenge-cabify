package car

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type CarDbImp struct {
	Cars map[int64]*model.Car
}

func (c CarDbImp) GetCarById(carID int64) *model.Car {
	return c.Cars[carID]
}

func (c CarDbImp) UpsertCar(car *model.Car) {
	c.Cars[car.Id] = car
}

func (c CarDbImp) GetAllCars() map[int64]*model.Car {
	return c.Cars
}

func (c CarDbImp) ResetCars() {
	c.Cars = make(map[int64]*model.Car)
}
