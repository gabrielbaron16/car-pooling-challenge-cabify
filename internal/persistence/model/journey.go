package model

type Journey struct {
	Id         int64 `json:"id"`
	Passengers uint  `json:"passengers"`
	AssignedTo *Car  `json:"assignedTo"`
}

type PendingOrderedMap struct {
	Journeys map[int64]*Journey
	Ids      []int64
}
