package models

import (
	"time"
)

type Post struct {
	Accounts     []*Account
	Message      string
	ScheduleTime time.Time
	Extras       map[string]interface{}
}
