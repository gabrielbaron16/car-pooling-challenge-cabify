package locate

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/mapper"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	journeyService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/journey"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/mocks"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestLocateHandlerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Locate Journey Handler Test Suite")
}

var _ = Describe("Locate Journey Handler Test Suite", func() {
	var (
		request            *http.Request
		mockCtrl           *gomock.Controller
		mockJourneyService *mocks.MockIJourneyService
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockJourneyService = mocks.NewMockIJourneyService(mockCtrl)
		journeyService.SetInstance(mockJourneyService)
		request, _ = http.NewRequest("POST", "/locate/{journey_id}", nil)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("PostLocateHandler function", func() {

		It("Response 200 - OK", func() {
			carResponse := getCarResponse()
			mockJourneyService.EXPECT().LocateJourney(int64(1)).Return(carResponse, nil)
			handlerResponse := PostLocateHandler(operations.PostLocateJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostLocateJourneyIDOK().WithPayload(mapper.MapCarToCarDTO(carResponse))))
		})

		It("Response 204 - No Content", func() {
			mockJourneyService.EXPECT().LocateJourney(int64(1)).Return(nil, nil)
			handlerResponse := PostLocateHandler(operations.PostLocateJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostLocateJourneyIDNoContent()))
		})

		It("Response 404 - Not Found", func() {
			mockJourneyService.EXPECT().LocateJourney(int64(1)).Return(nil, utils.ErrNotFound)
			handlerResponse := PostLocateHandler(operations.PostLocateJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostLocateJourneyIDNotFound()))
		})

		It("Response 500 - Internal Server Error", func() {
			mockJourneyService.EXPECT().LocateJourney(int64(1)).Return(nil, errors.New("error"))
			handlerResponse := PostLocateHandler(operations.PostLocateJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostLocateJourneyIDInternalServerError()))
		})
	})
})

func getCarResponse() *model.Car {
	return &model.Car{
		ID:             1,
		AvailableSeats: 4,
		MaxSeats:       4,
	}
}
