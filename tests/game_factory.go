package tests

import (
	"licorne/models"

	"gopkg.in/mgo.v2/bson"
)

func CreateFactoryGame() (map[string]interface{}, models.Game) {
	game_attr := map[string]interface{}{
		"name": "hello",
		"id": bson.NewObjectId(),
	}

	gameFactory := models.Game{
		Name: game_attr["name"].(string),
		ID: game_attr["id"].(bson.ObjectId),
	}
	return game_attr, gameFactory
}
