package journey

import (
	carDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/car"
	journeyDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/journey"
	pendingDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/pending"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"sync"
)

type JourneyService struct {
	Mu sync.Mutex
}

func (s *JourneyService) CreateJourney(journey *model.Journey) error {
	consultedJourney := journeyDb.GetInstance().GetJourneyById(journey.Id)
	if consultedJourney != nil {
		return utils.ErrDuplicatedID
	}
	cars := carDb.GetInstance().GetAllCars()
	s.assignCar(journey, cars)
	journeyDb.GetInstance().UpsertJourney(journey)
	return nil
}

func (s *JourneyService) LocateJourney(journeyID uint) (*model.Car, error) {
	journey := journeyDb.GetInstance().GetJourneyById(journeyID)
	if journey == nil {
		return nil, utils.ErrNotFound
	}
	var car *model.Car
	if journey.AssignedTo != nil {
		car = journey.AssignedTo
	}
	return car, nil
}

func (s *JourneyService) assignCar(journey *model.Journey, cars map[uint]*model.Car) {
	s.Mu.Lock()
	carAvailable := getCarAvailable(journey, cars)
	if carAvailable != nil {
		journey.AssignedTo = carAvailable
		carAvailable.AvailableSeats -= journey.Passengers
		carDb.GetInstance().UpsertCar(carAvailable)
	} else {
		pendingDb.GetInstance().AddPending(journey)
	}
	defer s.Mu.Unlock()
}

func getCarAvailable(journey *model.Journey, cars map[uint]*model.Car) *model.Car {
	var carAvailable *model.Car
	for _, car := range cars {
		if journey.Passengers <= car.AvailableSeats {
			carAvailable = car
			break
		}
	}
	return carAvailable
}
