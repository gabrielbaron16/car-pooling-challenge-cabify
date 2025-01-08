package dropoff

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
	"testing"
)

func TestDropoffServiceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dropoff Service Test Suite")
}

var _ = Describe("Dropoff Service Test Suite", func() {
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
		SetInstance(DropoffService{})
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Dropoff function", func() {

		It("Result Successfully - Journey Finished", func() {
			journey := getJourney()
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(journey)
			car := journey.AssignedTo
			mockJourneyDb.EXPECT().RemoveJourney(journey.Id).Times(1)
			car.AvailableSeats = car.AvailableSeats + journey.Passengers
			mockCarDb.EXPECT().UpsertCar(car).Times(1)
			car, err := GetInstance().Dropoff(journey.Id)
			Expect(err).To(BeNil())
			Expect(car).To(Equal(car))
		})

		It("Result Successfully - Journey removed from pending", func() {
			journey := getJourney()
			journey.AssignedTo = nil
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(journey)
			mockJourneyDb.EXPECT().RemoveJourney(journey.Id).Times(1)
			mockPendingDb.EXPECT().RemovePending(journey.Id).Times(1)
			car, err := GetInstance().Dropoff(journey.Id)
			Expect(err).To(BeNil())
			Expect(car).To(BeNil())
		})

		It("Result Failed - Journey not found", func() {
			journey := getJourney()
			mockJourneyDb.EXPECT().GetJourneyById(journey.Id).Return(nil)
			car, err := GetInstance().Dropoff(journey.Id)
			Expect(err).To(Equal(utils.ErrNotFound))
			Expect(car).To(BeNil())
		})
	})
})

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
