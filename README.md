# Car Pooling Service

## Introduction 

This service implements a car pooling service. 

The car Pooling service implements a very simple API that can be used to track the assignment of cars to Journeys according to the seat capacity of the Cars and the number of people that will travel in each journey.

Cars can have 4, 5 or 6 seats.

People request Journeys in groups of 1 to 6. People in the same group want to ride on the same car. You can take any group at any car that has enough empty seats for them, no matter their current location. If it's not possible to accommodate them, they're willing to wait until there is a car available for them. Once a car is available for a group that is waiting, they should ride.

Once they get a car assigned, they will journey until the drop off, you cannot ask them to take another car (i.e. you cannot swap them to another car to make space for another group).

In terms of fairness of trip order: groups should be served as fast as possible, but the arrival order should be kept when possible. If group B arrives later than group A, it can only be served before group A if no car can serve group A.

For example: a group of 6 is waiting for a car and there are 4 empty seats at a car for 6; if a group of 2 requests a car you may take them in the car. This may mean that the group of 6 waits a long time, possibly until they become frustrated and leave.

## API

The interface provided by the service is a RESTfull API. The operations are as follows.

### GET /status

Indicate the service has started up correctly and is ready to accept requests.

Responses:

* **200 OK** When the service is ready to receive requests.

### PUT /Cars

Load the list of available Cars in the service and remove all previous data (existing Journeys and Cars). This method may be called more than once during the life cycle of the service.

**Body** _required_ The list of Cars to load.

**Content Type** `application/json`

Sample:

```json
[
  {
    "id": 1,
    "seats": 4
  },
  {
    "id": 2,
    "seats": 6
  }
]
```

Responses:

* **200 OK** When the list is registered correctly.
* **400 Bad Request** When there is a failure in the request format, expected headers, or the payload can't be unmarshalled.

### POST /journey

A group of people requests to perform a journey.

**Body** _required_ The group of people that wants to perform the journey

**Content Type** `application/json`

Sample:

```json
{
  "id": 1,
  "people": 4
}
```

Responses:

* **200 OK** or **202 Accepted** When the group is registered correctly
* **400 Bad Request** When there is a failure in the request format or the payload can't be unmarshalled.

### POST /dropoff

A group of people requests to be dropped off. Whether they traveled or not.

**Body** _required_ A form with the group ID, such that `ID=X`

**Content Type** `application/x-www-form-urlencoded`

Responses:

* **200 OK** or **204 No Content** When the group is unregistered correctly.
* **404 Not Found** When the group is not to be found.
* **400 Bad Request** When there is a failure in the request format or the payload can't be unmarshalled.

### POST /locate

Given a group ID such that `ID=X`, return the car the group is traveling
with, or no car if they are still waiting to be served.

**Body** _required_ A url encoded form with the group ID such that `ID=X`

**Content Type** `application/x-www-form-urlencoded`

**Accept** `application/json`

Responses:

* **200 OK** With the car as the payload when the group is assigned to a car.
* **204 No Content** When the group is waiting to be assigned to a car.
* **404 Not Found** When the group is not to be found.
* **400 Bad Request** When there is a failure in the request format or the payload can't be unmarshalled.

## Swagger generation

To generate the swagger documentation, install go-swagger:

* apt-get update && apt-get install -y wget
* wget https://github.com/go-swagger/go-swagger/releases/download/v0.30.5/swagger_linux_amd64 -O /usr/local/bin/swagger
* chmod +x /usr/local/bin/swagger

Then, run the following command:

* swagger generate server -f ./swagger.yml --exclude-main

## Unit Tests

To run the tests, is necessary to generate the mocks, we use the mockgen tool from the gomock package. Follow these steps to generate the mocks:  
Install mockgen: If you haven't already installed mockgen, you can do so by running:  
go install github.com/golang/mock/mockgen@v1.6.0

Generate Mocks:

* mockgen -package mocks -source internal/service/car/car_service_interface.go -destination mocks/car_service_interface.go
* mockgen -package mocks -source internal/service/journey/journey_service_interface.go -destination mocks/journey_service_interface.go
* mockgen -package mocks -source internal/service/reassign/reassign_service_interface.go -destination mocks/reassign_service_interface.go
* mockgen -package mocks -source internal/service/dropoff/dropoff_service_interface.go -destination mocks/dropoff_service_interface.go
* mockgen -package mocks -source internal/persistence/db/car/db_car_interface.go -destination mocks/db_car_interface.go
* mockgen -package mocks -source internal/persistence/db/journey/db_journey_interface.go -destination mocks/db_journey_interface.go
* mockgen -package mocks -source internal/persistence/db/pending/db_pending_interface.go -destination mocks/db_pending_interface.go