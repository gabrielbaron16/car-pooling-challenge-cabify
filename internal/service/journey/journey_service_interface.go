package journey

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type IJourneyService interface {
	CreateJourney(journey *model.Journey) error
	LocateJourney(journeyId uint) (*model.Car, error)
}

var instance IJourneyService

func GetInstance() IJourneyService {
	return instance
}

func SetInstance(service IJourneyService) {
	instance = service
}
