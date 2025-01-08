package pending

import (
	"gitlab-hiring.cabify.tech/cabify/interviewing/car-pooling-challenge-go/internal/persistence/model"
)

type PendingDbImp struct {
	Pending *model.PendingOrderedMap
}

func (db PendingDbImp) AddPending(pending *model.Journey) {
	db.Pending.Journeys[pending.Id] = pending
	db.Pending.Ids = append(db.Pending.Ids, pending.Id)
}

func (db PendingDbImp) ResetPending() {
	db.Pending.Journeys = make(map[uint]*model.Journey)
	db.Pending.Ids = make([]uint, 0)
}

func (db PendingDbImp) RemovePending(journeyID uint) {
	delete(db.Pending.Journeys, journeyID)
	for i, p := range db.Pending.Ids {
		if p == journeyID {
			db.Pending.Ids = append(db.Pending.Ids[:i], db.Pending.Ids[i+1:]...)
			break
		}
	}
}

func (db PendingDbImp) GetAllPending() *model.PendingOrderedMap {
	return db.Pending
}
