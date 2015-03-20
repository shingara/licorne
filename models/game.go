package models

import (
	"errors"

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

// Generate the Game json
func MakeJsonGame(game Game) (map[string]interface{}) {
	json := map[string]interface{}{}
	json["data"] = ConvertJsonGame(game)
	return json
}

// Generate the Game list json
func MakeJsonGames(gameList []Game) (map[string]interface{}) {
	json := map[string]interface{}{}
	list := make([]GameJson, len(gameList))
	for i, game := range gameList {
		list[i] = ConvertJsonGame(game)
	}
	json["data"] = list
	return json
}

func (g Game) Valid() []error {
	errs := []error{}
	if (len(g.Players) < 2) {
		errs = append(errs, errors.New("need at least 2 players"))
	}

	if (len(g.Players) > 4) {
		errs = append(errs, errors.New("need less than 4 players"))
	}
	return errs
}
