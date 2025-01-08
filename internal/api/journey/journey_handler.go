package journey

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/mapper"
	journeyService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/journey"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
)

func PostJourneyHandler(params operations.PostJourneyParams) middleware.Responder {

	journeyModel := mapper.MapJourneyDTOToJourney(params.Journey)
	if err := journeyService.GetInstance().CreateJourney(journeyModel); err != nil {
		switch {
		case errors.Is(err, utils.ErrDuplicatedID):
			return operations.NewPostJourneyBadRequest()
		default:
			return operations.NewPostJourneyInternalServerError()
		}
	}
	return operations.NewPostJourneyAccepted()
}
