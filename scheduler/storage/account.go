package storage

import (
	"github.com/asdine/storm"
	"github.com/contentplanner/scheduler/scheduler/models"
)

type AccountStorage struct {
	DB storm.Node
}

func (storage *AccountStorage) Add(account *models.Account) error {
	err := storage.DB.Save(account)
	if err != nil {
		return err
	}
	return nil
}

func (storage *AccountStorage) One(ID int) (*models.Account, error) {
	var account models.Account
	err := storage.DB.One("ID", ID, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (storage *AccountStorage) All() ([]models.Account, error) {
	var accounts []models.Account
	err := storage.DB.All(&accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
