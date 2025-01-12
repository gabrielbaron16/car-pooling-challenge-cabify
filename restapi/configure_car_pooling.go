package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/api/car"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/api/dropoff"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/api/journey"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/api/locate"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
	"net/http"
)

func configureFlags(api *operations.CarPoolingAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CarPoolingAPI) http.Handler {
	// configure the api here
	errors.DefaultHTTPCode = http.StatusBadRequest
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.GetHandler = operations.GetHandlerFunc(func(params operations.GetParams) middleware.Responder {
		return operations.NewGetStatusOK().WithPayload("status: ok")
	})

	api.GetStatusHandler = operations.GetStatusHandlerFunc(func(params operations.GetStatusParams) middleware.Responder {
		return operations.NewGetStatusOK().WithPayload("status: ok")
	})

	api.PutCarsHandler = operations.PutCarsHandlerFunc(func(params operations.PutCarsParams) middleware.Responder {
		return car.PutCarsHandler(params)
	})

	api.PostJourneyHandler = operations.PostJourneyHandlerFunc(func(params operations.PostJourneyParams) middleware.Responder {
		return journey.PostJourneyHandler(params)
	})

	api.PostLocateJourneyIDHandler = operations.PostLocateJourneyIDHandlerFunc(func(params operations.PostLocateJourneyIDParams) middleware.Responder {
		return locate.PostLocateHandler(params)
	})

	api.PostDropoffJourneyIDHandler = operations.PostDropoffJourneyIDHandlerFunc(func(params operations.PostDropoffJourneyIDParams) middleware.Responder {
		return dropoff.PostDropoffHandler(params)
	})

	// Implement other handlers here

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
