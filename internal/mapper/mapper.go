package mapper

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/models"
)

func MapCarsDTOToCars(cars []*models.Car) []*model.Car {
	var result []*model.Car
	for _, car := range cars {
		result = append(result, &model.Car{
			ID:             *car.ID,
			MaxSeats:       uint(*car.Seats),
			AvailableSeats: uint(*car.Seats),
		})
	}
	return result
}

func MapCarToCarDTO(car *model.Car) *models.Car {
	seats := int32(car.MaxSeats)
	return &models.Car{
		ID:    &car.ID,
		Seats: &seats,
	}
}

func MapJourneyDTOToJourney(journey *models.Journey) *model.Journey {
	return &model.Journey{
		Id:         *journey.ID,
		Passengers: uint(*journey.People),
	}
}
