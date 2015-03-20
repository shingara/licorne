package tests

import (
	"licorne/models"

	"gopkg.in/mgo.v2/bson"
)

func CreateFactoryGame() (map[string]interface{}, models.Game) {
	players_attr := []models.Player{}
	players_attr = append(players_attr, models.Player{Name: "cyril", Color:"blue", ID: bson.NewObjectId()})
	players_attr = append(players_attr, models.Player{Name: "lena", Color:"pink", ID: bson.NewObjectId()})
	game_attr := map[string]interface{}{
		"name": "hello",
		"id": bson.NewObjectId(),
		"players":  players_attr,
	}

	gameFactory := models.Game{
		Name: game_attr["name"].(string),
		ID: game_attr["id"].(bson.ObjectId),
		Players: game_attr["players"].([]models.Player),
	}
	return game_attr, gameFactory
}
