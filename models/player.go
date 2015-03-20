package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Player struct {
		ID  bson.ObjectId `bson:"_id" json:"id"`
		Name string `json:"name" schema:"name"`
		Color string `json:"color" schema:"color"`
	}
)
