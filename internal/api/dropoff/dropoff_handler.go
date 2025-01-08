package dropoff

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	dropoffService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/dropoff"
	reassignService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/reassign"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
)

func PostDropoffHandler(params operations.PostDropoffJourneyIDParams) middleware.Responder {

	car, err := dropoffService.GetInstance().Dropoff(uint(params.JourneyID))
	if errors.Is(err, utils.ErrNotFound) {
		return operations.NewPostDropoffJourneyIDNotFound()
	}
	if err != nil {
		return operations.NewPostDropoffJourneyIDInternalServerError()
	}
	err = reassignService.GetInstance().Reassign(car)
	if err != nil {
		return operations.NewPostDropoffJourneyIDInternalServerError()
	}
	return operations.NewPostDropoffJourneyIDNoContent()
}
