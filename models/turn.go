package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)


type(

	PlayerStatus struct {
		PlayerId bson.ObjectId `bson:"_id" json:"id"`
		Position int
		NbDiamond int
	}

	Turn struct {
		GameId bson.ObjectId `bson:"_id" json:"id"`
		Number int

		StartTurn time.Time
		PlayerTurn bson.ObjectId `bson:"_id" json:"player_turn"`
		EndTurn time.Time

		PlayersEndStatus []PlayerStatus
	}
)
