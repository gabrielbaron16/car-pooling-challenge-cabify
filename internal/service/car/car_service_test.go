package car

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

func TestCarServiceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Car Service Test Suite")
}

var _ = Describe("Car Service Test Suite", func() {
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
		SetInstance(CarService{})
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("ResetCars function", func() {

		It("Result Successfully", func() {
			cars := getCarsArray()
			mockCarDb.EXPECT().ResetCars().Times(1)
			mockJourneyDb.EXPECT().ResetJourneys().Times(1)
			mockPendingDb.EXPECT().ResetPending().Times(1)
			mockCarDb.EXPECT().GetCarById(cars[0].ID).Return(nil)
			mockCarDb.EXPECT().UpsertCar(cars[0]).Times(1)
			err := GetInstance().ResetCars(cars)
			Expect(err).To(BeNil())
		})

		It("Result Failed - Duplicated Id", func() {
			cars := getCarsArrayDuplicated()
			mockCarDb.EXPECT().ResetCars().Times(2)
			mockJourneyDb.EXPECT().ResetJourneys().Times(1)
			mockPendingDb.EXPECT().ResetPending().Times(1)
			mockCarDb.EXPECT().GetCarById(cars[0].ID).Return(nil)
			mockCarDb.EXPECT().UpsertCar(cars[0]).Times(1)
			mockCarDb.EXPECT().GetCarById(cars[1].ID).Return(cars[0])
			err := GetInstance().ResetCars(cars)
			Expect(err).To(Equal(utils.ErrDuplicatedID))
		})
	})
})

func getCarsArray() []*model.Car {
	return []*model.Car{
		{
			ID:             1,
			AvailableSeats: 4,
			MaxSeats:       4,
		},
	}
}

func getCarsArrayDuplicated() []*model.Car {
	return []*model.Car{
		{
			ID:             1,
			AvailableSeats: 4,
			MaxSeats:       4,
		},
		{
			ID:             1,
			AvailableSeats: 6,
			MaxSeats:       6,
		},
	}
}
