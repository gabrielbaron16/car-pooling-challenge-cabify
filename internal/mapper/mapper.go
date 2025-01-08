package mapper

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/models"
)

func MapCarsDTOToCars(cars []*models.Car) []*model.Car {
	var result []*model.Car
	for _, car := range cars {
		result = append(result, &model.Car{
			ID:             uint(car.ID),
			MaxSeats:       uint(car.Seats),
			AvailableSeats: uint(car.Seats),
		})
	}
	return result
}

func MapCarToCarDTO(car *model.Car) *models.Car {
	return &models.Car{
		ID:    int64(car.ID),
		Seats: int32(car.AvailableSeats),
	}
}

func MapJourneyDTOToJourney(journey *models.Journey) *model.Journey {
	return &model.Journey{
		Id:         uint(journey.ID),
		Passengers: uint(journey.People),
	}
}
