package pending

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

var pendingDAO PendingDatabase

type PendingDatabase interface {
	AddPending(pending *model.Journey)
	ResetPending()
	RemovePending(journeyID int64)
	GetAllPending() *model.PendingOrderedMap
}

func SetInstance(instance PendingDatabase) {
	pendingDAO = instance
}

func GetInstance() PendingDatabase {
	return pendingDAO
}
