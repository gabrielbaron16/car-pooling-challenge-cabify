package dropoff

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
	dropoffService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/dropoff"
	reassignService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/reassign"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/mocks"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestDropoffHandlerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Locate Journey Handler Test Suite")
}

var _ = Describe("Dropoff Handler Test Suite", func() {
	var (
		request             *http.Request
		mockCtrl            *gomock.Controller
		mockDropoffService  *mocks.MockIDropoffService
		mockReassignService *mocks.MockIReassignService
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockDropoffService = mocks.NewMockIDropoffService(mockCtrl)
		dropoffService.SetInstance(mockDropoffService)
		mockReassignService = mocks.NewMockIReassignService(mockCtrl)
		reassignService.SetInstance(mockReassignService)
		request, _ = http.NewRequest("POST", "/dropoff/{journey_id}", nil)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("PostDropoffHandler function", func() {

		It("Response 204 - No Content", func() {
			carResponse := getCarResponse()
			mockDropoffService.EXPECT().Dropoff(uint(1)).Return(carResponse, nil)
			mockReassignService.EXPECT().Reassign(carResponse).Return(nil)
			handlerResponse := PostDropoffHandler(operations.PostDropoffJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostDropoffJourneyIDNoContent()))
		})

		It("Response 404 - Not Found", func() {
			mockDropoffService.EXPECT().Dropoff(uint(1)).Return(nil, utils.ErrNotFound)
			handlerResponse := PostDropoffHandler(operations.PostDropoffJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostDropoffJourneyIDNotFound()))
		})

		It("Response 500 - Internal Server Error (Error on dropoff process)", func() {
			mockDropoffService.EXPECT().Dropoff(uint(1)).Return(nil, errors.New("error"))
			handlerResponse := PostDropoffHandler(operations.PostDropoffJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostDropoffJourneyIDInternalServerError()))
		})

		It("Response 500 - Internal Server Error (Error on reassign process)", func() {
			carResponse := getCarResponse()
			mockDropoffService.EXPECT().Dropoff(uint(1)).Return(carResponse, nil)
			mockReassignService.EXPECT().Reassign(carResponse).Return(errors.New("error"))
			handlerResponse := PostDropoffHandler(operations.PostDropoffJourneyIDParams{
				HTTPRequest: request,
				JourneyID:   1,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostDropoffJourneyIDInternalServerError()))
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
