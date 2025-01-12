package reassign

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type IReassignService interface {
	Reassign(car *model.Car)
}

var instance IReassignService

func GetInstance() IReassignService {
	return instance
}

func SetInstance(service IReassignService) {
	instance = service
}
