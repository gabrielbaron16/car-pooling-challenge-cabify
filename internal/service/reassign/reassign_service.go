package reassign

import (
	"fmt"
	carDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/car"
	journeyDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/journey"
	pendingDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/pending"
	model2 "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"sync"
)

type ReassignService struct {
	Mu sync.Mutex
}

func (s *ReassignService) Reassign(car *model2.Car) {
	s.Mu.Lock()
	if car != nil {
		pending := pendingDb.GetInstance().GetAllPending()
		nextJourney := getNextJourney(pending.Ids, pending.Journeys, car.AvailableSeats)
		if nextJourney != nil {
			fmt.Printf(">> Car %d reassigned to journey %d\n", car.Id, nextJourney.Id)
			pending.Journeys[nextJourney.Id].AssignedTo = car
			updatedJourney := pending.Journeys[nextJourney.Id]
			car.AvailableSeats -= pending.Journeys[nextJourney.Id].Passengers
			journeyDb.GetInstance().UpsertJourney(updatedJourney)
			carDb.GetInstance().UpsertCar(car)
			pendingDb.GetInstance().RemovePending(nextJourney.Id)
		}
	}
	defer s.Mu.Unlock()
}

func getNextJourney(ids []int64, journeys map[int64]*model2.Journey, availableSeats uint) *model2.Journey {
	for _, id := range ids {
		if journeys[id].Passengers <= availableSeats {
			return journeys[id]
		}
	}
	return nil
}
