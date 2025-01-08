package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	injection "gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/dependency-injection"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi"
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/restapi/operations"
)

func main() {

	injection.InjectCarpoolServeDependencies()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCarPoolingAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Car Pooling"
	parser.LongDescription = "This is a sample server for car pooling."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
