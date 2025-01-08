package car

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/mapper"
	carService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/car"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
)

func PutCarsHandler(params operations.PutCarsParams) middleware.Responder {
	carsArray := mapper.MapCarsDTOToCars(params.Cars)

	if err := carService.GetInstance().ResetCars(carsArray); err != nil {
		switch {
		case errors.Is(err, utils.ErrDuplicatedID):
			return operations.NewPutCarsBadRequest()
		default:
			return operations.NewPutCarsInternalServerError()
		}
	}
	return operations.NewPutCarsOK()
}
