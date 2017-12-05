package storage

import (
	"fmt"

	"github.com/contentplanner/scheduler/scheduler/models"
)

type PostStorage struct{}

func (storage *PostStorage) Schedule(post *models.Post) []error {
	var errors []error

	for _, account := range post.Accounts {
		if valid := account.Media.Validate(); !valid {
			errors = append(errors, fmt.Errorf("Media credentials not valid"))
			continue
		}

		err := account.Media.Post(post)
		if err != nil {
			errors = append(errors, err)
			continue
		}
	}

	return errors
}
