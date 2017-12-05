package common

type credentials struct {
	Facebook struct {
		AppID     string
		AppSecret string
	}
}

var Credentials credentials

func init() {
	Credentials = credentials{}
}
