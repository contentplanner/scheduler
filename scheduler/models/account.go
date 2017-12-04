package models

import (
	"encoding/json"
)

type Account struct {
	ID    int    `storm:"id,increment"`
	Name  string `storm:"index"`
	Media Media
}

func (ac *Account) UnmarshalJSON(b []byte) error {
	var objMap map[string]interface{}
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	// ID and Name
	ac.ID = int(objMap["ID"].(float64))
	ac.Name = objMap["Name"].(string)

	// media
	mediaMap := objMap["Media"].(map[string]interface{})
	mediaMapBytes, err := json.Marshal(mediaMap)
	if err != nil {
		return err
	}

	// actual media type
	var actual Media
	switch mediaMap["type"] {
	case "facebook":
		actual = &Facebook{}
		break
	}

	// unmarshal media
	err = json.Unmarshal(mediaMapBytes, actual)
	if err != nil {
		return err
	}

	// set Media
	ac.Media = actual

	return nil
}
