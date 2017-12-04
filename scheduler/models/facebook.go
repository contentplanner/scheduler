package models

import (
	"encoding/json"

	"github.com/contentplanner/scheduler/scheduler/common"
	"github.com/huandu/facebook"
)

type Facebook struct {
	ID    string
	Token string
}

func (fb *Facebook) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["type"] = "facebook"
	m["id"] = fb.ID
	m["token"] = fb.Token
	return json.Marshal(m)
}

func (fb *Facebook) UnmarshalJSON(b []byte) error {
	m := make(map[string]string)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	fb.ID = m["id"]
	fb.Token = m["token"]
	return nil
}

func (fb *Facebook) Validate(creds common.Credentials) bool {
	return creds.Facebook.AppID != "" && creds.Facebook.AppSecret != ""
}

func (fb *Facebook) Post(post *Post, creds common.Credentials) error {
	fbApp := facebook.New(creds.Facebook.AppID, creds.Facebook.AppSecret)

	// validate fb session
	session := fbApp.Session(fb.Token)
	err := session.Validate()
	if err != nil {
		return err
	}

	// Params
	params := facebook.Params{}

	params["published"] = false
	params["message"] = post.Message
	params["scheduled_publish_time"] = post.ScheduleTime.Unix()

	if link, ok := post.Extras["link"]; ok {
		params["link"] = link
	}

	// Post to feed
	_, err = session.Post("/"+fb.ID+"/feed", params)
	return err
}
