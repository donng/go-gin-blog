package model

type Tag struct {
	Model
	Name string `json:"name"`
	Status int `json:"status"`
}
