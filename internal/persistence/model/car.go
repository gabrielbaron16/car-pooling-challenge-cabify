package model

type Car struct {
	Id             int64 `json:"id"`
	MaxSeats       uint  `json:"maxSeats"`
	AvailableSeats uint  `json:"availableSeats"`
}
