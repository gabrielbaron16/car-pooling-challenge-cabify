package journey

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type JourneyDbImp struct {
	Journeys map[uint]*model.Journey
}

func (c JourneyDbImp) GetJourneyById(journeyID uint) *model.Journey {
	return c.Journeys[journeyID]
}

func (c JourneyDbImp) UpsertJourney(journey *model.Journey) {
	c.Journeys[journey.Id] = journey
}

func (c JourneyDbImp) ResetJourneys() {
	c.Journeys = make(map[uint]*model.Journey)
}

func (c JourneyDbImp) RemoveJourney(journeyID uint) {
	delete(c.Journeys, journeyID)
}
