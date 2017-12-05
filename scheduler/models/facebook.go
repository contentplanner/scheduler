package models

import (
	"encoding/json"

	"github.com/contentplanner/scheduler/scheduler/common"
	"github.com/huandu/facebook"
)

type (
	Facebook struct {
		ID    string
		Token string
	}
	facebookPost struct {
		ID          string `facebook:"id"`
		PublishTime string `facebook:"scheduled_publish_time"`
		Message     string `facebook:"message"`
		Type        string `facebook:"type"`
	}
)

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

func (fb *Facebook) Validate() bool {
	return common.Credentials.Facebook.AppID != "" && common.Credentials.Facebook.AppSecret != ""
}

func (fb *Facebook) Post(post *Post) error {
	fbApp := facebook.New(common.Credentials.Facebook.AppID, common.Credentials.Facebook.AppSecret)

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

func (fb *Facebook) GetPosts(page int) (interface{}, error) {
	fbApp := facebook.New(common.Credentials.Facebook.AppID, common.Credentials.Facebook.AppSecret)

	// validate fb session
	session := fbApp.Session(fb.Token)
	err := session.Validate()
	if err != nil {
		return nil, err
	}

	res, err := session.Get("/"+fb.ID+"/scheduled_posts", facebook.Params{
		"fields": "id,scheduled_publish_time,message,attachments,type",
		"limit":  100,
		"offset": 100 * (page - 1),
	})
	if err != nil {
		return nil, err
	}

	var posts []facebookPost
	err = res.DecodeField("data", &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
