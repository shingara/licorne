package models

type (
	Player struct {
		Name string `json:"name" schema:"name"`
		Color string `json:"color" schema:"color"`
	}
)
