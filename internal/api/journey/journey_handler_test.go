package journey

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/mapper"
	journeyService "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/service/journey"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/utils"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/mocks"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/models"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestJourneyHandlerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Journey Handler Test Suite")
}

var _ = Describe("Journey Handler Test Suite", func() {
	var (
		request            *http.Request
		mockCtrl           *gomock.Controller
		mockJourneyService *mocks.MockIJourneyService
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockJourneyService = mocks.NewMockIJourneyService(mockCtrl)
		journeyService.SetInstance(mockJourneyService)
		request, _ = http.NewRequest("POST", "/journey", nil)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("PostJourneyHandler function", func() {

		It("Response 202 - Accepted", func() {
			journey := getPostJourneyBody()
			mockJourneyService.EXPECT().CreateJourney(mapper.MapJourneyDTOToJourney(journey)).Return(nil)
			handlerResponse := PostJourneyHandler(operations.PostJourneyParams{
				HTTPRequest: request,
				Journey:     journey,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostJourneyAccepted()))
		})

		It("Response 400 - Bad Request", func() {
			journey := getPostJourneyBody()
			mockJourneyService.EXPECT().CreateJourney(mapper.MapJourneyDTOToJourney(journey)).Return(utils.ErrDuplicatedID)
			handlerResponse := PostJourneyHandler(operations.PostJourneyParams{
				HTTPRequest: request,
				Journey:     journey,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostJourneyBadRequest()))
		})

		It("Response 500 - Internal Server Error", func() {
			journey := getPostJourneyBody()
			mockJourneyService.EXPECT().CreateJourney(mapper.MapJourneyDTOToJourney(journey)).Return(errors.New("error"))
			handlerResponse := PostJourneyHandler(operations.PostJourneyParams{
				HTTPRequest: request,
				Journey:     journey,
			})
			Expect(handlerResponse).To(BeEquivalentTo(operations.NewPostJourneyInternalServerError()))
		})
	})
})

func getPostJourneyBody() *models.Journey {
	id := int64(1)
	people := int32(4)
	return &models.Journey{
		ID:     &id,
		People: &people,
	}
}
