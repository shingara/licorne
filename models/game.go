package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// GameModel contains information for a game
	Game struct {
		ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name string `json:"name" schema:"name"`
		Players []Player `json:"players" schema:"players"`
	}

	GameJson struct {
		Game
		Type string `json:"type", omitempty`
	}

	GameForm struct {
		Game
	}
)
