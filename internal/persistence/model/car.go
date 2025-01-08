package model

type Car struct {
	ID             int64 `json:"id"`
	MaxSeats       uint  `json:"maxSeats"`
	AvailableSeats uint  `json:"availableSeats"`
}
