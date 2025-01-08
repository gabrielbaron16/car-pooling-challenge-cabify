package dependency_injection

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/car"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/journey"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/pending"
	model2 "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	carService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/car"
	dropoffService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/dropoff"
	journeyService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/journey"
	reassignService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/reassign"
	"sync"
)

func InjectCarpoolServeDependencies() {

	// service
	carService.SetInstance(carService.CarService{})
	journeyService.SetInstance(&journeyService.JourneyService{
		Mu: sync.Mutex{},
	})
	dropoffService.SetInstance(dropoffService.DropoffService{})
	reassignService.SetInstance(&reassignService.ReassignService{
		Mu: sync.Mutex{},
	})

	// database
	car.SetInstance(&car.CarDbImp{
		Cars: make(map[uint]*model2.Car),
	})
	journey.SetInstance(&journey.JourneyDbImp{
		Journeys: make(map[uint]*model2.Journey),
	})
	pending.SetInstance(&pending.PendingDbImp{
		Pending: &model2.PendingOrderedMap{
			Journeys: map[uint]*model2.Journey{},
			Ids:      []uint{},
		},
	})
}
