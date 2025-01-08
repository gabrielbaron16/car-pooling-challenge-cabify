package dropoff

import (
	carDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/car"
	journeyDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/journey"
	pendingDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/pending"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
)

type DropoffService struct{}

func (s DropoffService) Dropoff(journeyID uint) (car *model.Car, err error) {
	journey := journeyDb.GetInstance().GetJourneyById(journeyID)
	if journey == nil {
		return nil, utils.ErrNotFound
	}

	car = journey.AssignedTo
	journeyDb.GetInstance().RemoveJourney(journeyID)

	if car != nil {
		car.AvailableSeats += journey.Passengers
		carDb.GetInstance().UpsertCar(car)
	} else {
		pendingDb.GetInstance().RemovePending(journeyID)
	}
	return car, nil
}
