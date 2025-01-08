package car

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/mapper"
	carService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/car"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/mocks"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/models"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestCarHandlerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Car Handler Test Suite")
}

var _ = Describe("Car Handler Test Suite", func() {
	var (
		request        *http.Request
		mockCtrl       *gomock.Controller
		mockCarService *mocks.MockICarService
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCarService = mocks.NewMockICarService(mockCtrl)
		carService.SetInstance(mockCarService)
		request, _ = http.NewRequest("PUT", "/cars", nil)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("PutCarsHandler function", func() {

		It("Response 200 - OK", func() {
			cars := getPutCarsBody()
			mockCarService.EXPECT().ResetCars(mapper.MapCarsDTOToCars(cars)).Return(nil)
			handlerResponse := PutCarsHandler(operations.PutCarsParams{
				HTTPRequest: request,
				Cars:        cars,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPutCarsOK()))
		})

		It("Response 400 - BadRequest", func() {
			cars := getPutCarsBody()
			mockCarService.EXPECT().ResetCars(mapper.MapCarsDTOToCars(cars)).Return(utils.ErrDuplicatedID)
			handlerResponse := PutCarsHandler(operations.PutCarsParams{
				HTTPRequest: request,
				Cars:        cars,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPutCarsOK()))
		})

		It("Response 500 - InternalServerError", func() {
			cars := getPutCarsBody()
			mockCarService.EXPECT().ResetCars(mapper.MapCarsDTOToCars(cars)).Return(errors.New("error"))
			handlerResponse := PutCarsHandler(operations.PutCarsParams{
				HTTPRequest: request,
				Cars:        cars,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPutCarsInternalServerError()))
		})
	})
})

func getPutCarsBody() []*models.Car {
	id1 := int64(1)
	seats1 := int32(4)
	id2 := int64(2)
	seats2 := int32(4)
	return []*models.Car{
		{
			ID:    &id1,
			Seats: &seats1,
		},
		{
			ID:    &id2,
			Seats: &seats2,
		},
	}
}
