package dropoff

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type IDropoffService interface {
	Dropoff(journeyID int64) (car *model.Car, err error)
}

var instance IDropoffService

func GetInstance() IDropoffService {
	return instance
}

func SetInstance(service IDropoffService) {
	instance = service
}
