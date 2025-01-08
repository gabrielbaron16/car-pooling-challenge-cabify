package journey

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

var journeyDAO JourneyDatabase

type JourneyDatabase interface {
	GetJourneyById(journeyID uint) *model.Journey
	UpsertJourney(journey *model.Journey)
	ResetJourneys()
	RemoveJourney(journeyID uint)
}

func GetInstance() JourneyDatabase {
	return journeyDAO
}

func SetInstance(instance JourneyDatabase) {
	journeyDAO = instance
}
