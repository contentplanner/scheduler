package models

import "github.com/contentplanner/scheduler/scheduler/common"

type Media interface {
	Validate(common.Credentials) bool
	Post(*Post, common.Credentials) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
