package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// GameModel contains information for a game
	Game struct {
		ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name string `json:"name"`
		Players []Player `json:"players"`
	}

	GameJson struct {
		ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name string `json:"name"`
		Type string `json:"type", omitempty`
		Players []Player `json:"players"`
	}
)
