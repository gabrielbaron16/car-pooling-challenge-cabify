package model

type Journey struct {
	Id         uint `json:"id"`
	Passengers uint `json:"passengers"`
	AssignedTo *Car `json:"assignedTo"`
}

type PendingOrderedMap struct {
	Journeys map[uint]*Journey
	Ids      []uint
}
