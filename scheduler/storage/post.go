package storage

import (
	"fmt"

	"github.com/contentplanner/scheduler/scheduler/common"
	"github.com/contentplanner/scheduler/scheduler/models"
)

type PostStorage struct {
	Credentials common.Credentials
}

func (storage *PostStorage) SetFacebook(appID string, appSecret string) {
	storage.Credentials.Facebook.AppID = appID
	storage.Credentials.Facebook.AppSecret = appSecret
}

func (storage *PostStorage) Schedule(post *models.Post) []error {
	var errors []error

	for _, account := range post.Accounts {
		if valid := account.Media.Validate(storage.Credentials); !valid {
			errors = append(errors, fmt.Errorf("Media credentials not valid"))
			continue
		}

		err := account.Media.Post(post, storage.Credentials)
		if err != nil {
			errors = append(errors, err)
			continue
		}
	}

	return errors
}
