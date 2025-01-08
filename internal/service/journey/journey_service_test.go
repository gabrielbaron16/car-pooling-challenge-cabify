package journey

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	carDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/car"
	journeyDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/journey"
	pendingDb "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/db/pending"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/mocks"
	"go.uber.org/mock/gomock"
	"sync"
	"testing"
)

func TestJourneyServiceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Journey Service Test Suite")
}

var _ = Describe("Journey Service Test Suite", func() {
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
		SetInstance(&JourneyService{
			Mu: sync.Mutex{},
		})
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Create Journey function", func() {

		It("Result Successfully - Car Assigned to Journey", func() {
			journey := getJourney()
			journey.AssignedTo = nil
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(nil)
			cars := getCarsMap()
			mockCarDb.EXPECT().GetAllCars().Return(cars)
			car := getCar()
			car.AvailableSeats = car.AvailableSeats - journey.Passengers
			journey.AssignedTo = car
			mockCarDb.EXPECT().UpsertCar(car).Times(1)
			mockJourneyDb.EXPECT().UpsertJourney(journey).Times(1)
			journeyToCreate := getJourney()
			err := GetInstance().CreateJourney(journeyToCreate)
			Expect(err).To(BeNil())
		})

		It("Result Successfully - Journey Added to pending", func() {
			journey := getJourney()
			journey.AssignedTo = nil
			journey.Passengers = 5
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(nil)
			cars := getCarsMap()
			mockCarDb.EXPECT().GetAllCars().Return(cars)
			mockPendingDb.EXPECT().AddPending(journey).Times(1)
			mockJourneyDb.EXPECT().UpsertJourney(journey).Times(1)
			err := GetInstance().CreateJourney(journey)
			Expect(err).To(BeNil())
		})

		It("Result Failed - Journey Duplicated id", func() {
			journey := getJourney()
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(journey)
			err := GetInstance().CreateJourney(journey)
			Expect(err).To(Equal(utils.ErrDuplicatedID))
		})
	})

	Context("Locate Journey function", func() {

		It("Result Successfully - Journey is in progress", func() {
			journey := getJourney()
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(journey)
			car, err := GetInstance().LocateJourney(journey.Id)
			Expect(err).To(BeNil())
			Expect(car).To(Equal(journey.AssignedTo))
		})

		It("Result Successfully - Journey is pending", func() {
			journey := getJourney()
			journey.AssignedTo = nil
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(journey)
			car, err := GetInstance().LocateJourney(journey.Id)
			Expect(err).To(BeNil())
			Expect(car).To(BeNil())
		})

		It("Result Failed - Journey not found", func() {
			journey := getJourney()
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(nil)
			car, err := GetInstance().LocateJourney(journey.Id)
			Expect(err).To(Equal(utils.ErrNotFound))
			Expect(car).To(BeNil())
		})
	})
})

func getCarsMap() map[int64]*model.Car {
	return map[int64]*model.Car{
		1: {
			ID:             1,
			AvailableSeats: 4,
			MaxSeats:       4,
		},
	}
}

func getCar() *model.Car {
	return &model.Car{
		ID:             1,
		AvailableSeats: 4,
		MaxSeats:       4,
	}
}

func getJourney() *model.Journey {
	return &model.Journey{
		Id:         1,
		Passengers: 4,
		AssignedTo: getCar(),
	}
}
