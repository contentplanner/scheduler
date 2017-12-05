package scheduler

import (
	"github.com/asdine/storm"
	"github.com/contentplanner/scheduler/scheduler/common"
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

func (scheduler *Scheduler) SetFacebook(appID string, appSecret string) {
	common.Credentials.Facebook.AppID = appID
	common.Credentials.Facebook.AppSecret = appSecret
}

func New(db *storm.DB) (*Scheduler, error) {
	scheduler := &Scheduler{
		db: db,
		AccountStorage: &storage.AccountStorage{
			DB: db.From("scheduler"),
		},
		PostStorage: &storage.PostStorage{},
	}
	return scheduler, nil
}

func Default() (*Scheduler, error) {
	db, err := storm.Open("scheduler.db")
	if err != nil {
		return nil, err
	}

	scheduler := &Scheduler{
		db: db,
		AccountStorage: &storage.AccountStorage{
			DB: db.From("scheduler"),
		},
		PostStorage: &storage.PostStorage{},
	}
	return scheduler, nil
}
