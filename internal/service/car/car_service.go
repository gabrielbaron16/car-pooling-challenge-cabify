package car

import (
	carDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/car"
	journeyDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/journey"
	pendingDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/pending"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
)

type CarService struct{}

func (s CarService) ResetCars(cars []*model.Car) error {
	carDb.GetInstance().ResetCars()
	journeyDb.GetInstance().ResetJourneys()
	pendingDb.GetInstance().ResetPending()
	for _, car := range cars {
		consultedCar := carDb.GetInstance().GetCarById(car.Id)
		if consultedCar != nil {
			carDb.GetInstance().ResetCars()
			return utils.ErrDuplicatedID
		}
		carDb.GetInstance().UpsertCar(car)
	}
	return nil
}
