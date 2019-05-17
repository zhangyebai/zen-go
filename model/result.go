package model

type (
	Page struct {
		Page  int `json:"page"`
		Size  int `json:"size"`
		Pages int `json:"pages"`
		Total int `json:"total"`
	}
)
