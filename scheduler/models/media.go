package models

type Media interface {
	Validate() bool
	Post(*Post) error
	GetPosts(int) (interface{}, error)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
