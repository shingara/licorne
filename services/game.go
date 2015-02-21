package services

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"licorne/models"
	"licorne/utilities"

)

var collection_name = "games"


func CreateGame(game *models.Game) (*models.Game, error) {
	if (!game.ID.Valid()) {
		game.ID = bson.NewObjectId()
	}
	err := utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.Insert(game)
	})
	return game, err
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
	bson_id := bson.ObjectIdHex(id.(string))
	err = utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.FindId(bson_id).One(&game)
	})
	return game, err
}
