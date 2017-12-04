package scheduler

import (
	"github.com/asdine/storm"
	"github.com/contentplanner/scheduler/scheduler/storage"
)

type Scheduler struct {
	db             *storm.DB
	AccountStorage *storage.AccountStorage
	PostStorage    *storage.PostStorage
}

func (scheduler *Scheduler) Close() error {
	return scheduler.db.Close()
}

func New() (*Scheduler, error) {
	db, err := storm.Open("scheduler.db")
	if err != nil {
		return nil, err
	}

	scheduler := &Scheduler{
		db: db,
		AccountStorage: &storage.AccountStorage{
			DB: db,
		},
		PostStorage: &storage.PostStorage{},
	}
	return scheduler, nil
}
