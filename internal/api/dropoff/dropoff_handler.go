package dropoff

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	dropoffService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/dropoff"
	reassignService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/reassign"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
)

func PostDropoffHandler(params operations.PostDropoffParams) middleware.Responder {

	car, err := dropoffService.GetInstance().Dropoff(params.ID)
	if errors.Is(err, utils.ErrNotFound) {
		return operations.NewPostDropoffNotFound()
	}
	if err != nil {
		return operations.NewPostDropoffInternalServerError()
	}
	reassignService.GetInstance().Reassign(car)
	return operations.NewPostDropoffNoContent()
}
