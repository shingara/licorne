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
		Game
		Type string `json:"type", omitempty`
	}

	GameForm struct {
		Name string `form:"name" binding:"required"`
		Players []Player `form:"players" binding:"required"`
	}
)

// Cast a GameForm to a Game model
func ConvertToGame(form GameForm) Game {
	return Game{
		Name: form.Name,
		Players: form.Players,
	}
}

func ConvertJsonGame(game Game) (GameJson) {
	return GameJson{
		Game: game,
		Type: "games",
	}
}

func MakeJsonGames(gameList []Game) (map[string]interface{}) {
	json := map[string]interface{}{}
	list := make([]GameJson, len(gameList))
	for i, game := range gameList {
		list[i] = ConvertJsonGame(game)
	}
	json["data"] = list
	return json
}
