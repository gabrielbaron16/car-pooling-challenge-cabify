package reassign

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	carDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/car"
	journeyDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/journey"
	pendingDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/pending"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/mocks"
	"go.uber.org/mock/gomock"
	"sync"
	"testing"
)

func TestReassignServiceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reassign Service Test Suite")
}

var _ = Describe("Reassign Service Test Suite", func() {
	var (
		mockCtrl      *gomock.Controller
		mockCarDb     *mocks.MockCarDatabase
		mockJourneyDb *mocks.MockJourneyDatabase
		mockPendingDb *mocks.MockPendingDatabase
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCarDb = mocks.NewMockCarDatabase(mockCtrl)
		carDb.SetInstance(mockCarDb)
		mockJourneyDb = mocks.NewMockJourneyDatabase(mockCtrl)
		journeyDb.SetInstance(mockJourneyDb)
		mockPendingDb = mocks.NewMockPendingDatabase(mockCtrl)
		pendingDb.SetInstance(mockPendingDb)
		SetInstance(&ReassignService{
			Mu: sync.Mutex{},
		})
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Reassign function", func() {

		It("Result Successfully - Car Reassigned to Journey", func() {
			car := getCar()
			pending := getPending()
			mockPendingDb.EXPECT().GetAllPending().Return(pending)
			updatedJourney := pending.Journeys[1]
			car.AvailableSeats = car.AvailableSeats - updatedJourney.Passengers
			updatedJourney.AssignedTo = car
			mockJourneyDb.EXPECT().UpsertJourney(updatedJourney).Times(1)
			mockCarDb.EXPECT().UpsertCar(car).Times(1)
			mockPendingDb.EXPECT().RemovePending(updatedJourney.Id).Times(1)
			carToReassign := getCar()
			GetInstance().Reassign(carToReassign)
		})

		It("Result Successfully - Car Not Reassigned", func() {
			car := getCar()
			pending := getPending()
			pending.Journeys[1].Passengers = 5
			mockPendingDb.EXPECT().GetAllPending().Return(pending)
			GetInstance().Reassign(car)
		})
	})
})

func getCar() *model.Car {
	return &model.Car{
		Id:             1,
		AvailableSeats: 4,
		MaxSeats:       4,
	}
}

func getPending() *model.PendingOrderedMap {
	return &model.PendingOrderedMap{
		Journeys: getJourneyMap(),
		Ids:      []int64{1},
	}
}

func getJourneyMap() map[int64]*model.Journey {
	return map[int64]*model.Journey{
		1: {
			Id:         1,
			Passengers: 4,
			AssignedTo: getCar(),
		},
	}
}
