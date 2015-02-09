package services

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"licorne/models"
	"licorne/utilities"

)

var collection_name = "games"


func CreateGame(game *models.Game) error {
	err := utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.Insert(game)
	})
	return err
}

func AllGame() (gameList []models.Game, err error) {
	gameList = []models.Game{}
	err = utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.Find(bson.M{}).All(&gameList)
	})
	return gameList, err
}

func GetGame(id interface{}) (game models.Game, err error) {
	game = models.Game{}
	err = utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.FindId(id).One(&game)
	})
	return game, err
}

func MakeJsonGames(gameList []models.Game) (map[string]interface{}) {
	json := map[string]interface{}{}
	list := make([]models.GameJson, len(gameList))
	for i, game := range gameList {
		json_game := models.GameJson{
			ID: game.ID,
			Name: game.Name,
			Type: "games",
			Players: game.Players,
		}
		list[i] = json_game
	}
	json["data"] = list
	return json
}
