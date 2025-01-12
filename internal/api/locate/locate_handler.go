package locate

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/mapper"
	journeyService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/journey"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
)

func PostLocateHandler(params operations.PostLocateParams) middleware.Responder {

	car, err := journeyService.GetInstance().LocateJourney(params.ID)
	if errors.Is(err, utils.ErrNotFound) {
		return operations.NewPostLocateJourneyIDNotFound()
	}
	if err != nil {
		return operations.NewPostLocateJourneyIDInternalServerError()
	}
	if car == nil {
		return operations.NewPostDropoffJourneyIDNoContent()
	}
	return operations.NewPostLocateJourneyIDOK().WithPayload(mapper.MapCarToCarDTO(car))
}
