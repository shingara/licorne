package services

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"licorne/models"
	"licorne/utilities"

)

var collection_name = "games"


// Save the game pass in args
func CreateGame(game *models.Game) (*models.Game, []error) {
	if (!game.ID.Valid()) {
		game.ID = bson.NewObjectId()
	}
	model_error := game.Valid()
	if (len(model_error) > 0) {
		return game, model_error
	}
	err := utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.Insert(game)
	})
	if (err != nil) {
		model_error = append(model_error, err)
	}
	return game, model_error
}

// Return all games save in DB
func AllGame() (gameList []models.Game, err error) {
	gameList = []models.Game{}
	err = utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.Find(bson.M{}).All(&gameList)
	})
	return gameList, err
}

// Return the game with the id pass
func GetGame(id interface{}) (game models.Game, err error) {
	game = models.Game{}
	bson_id := bson.ObjectIdHex(id.(string))
	err = utilities.WithCollection( collection_name, func(collection *mgo.Collection) error {
		return collection.FindId(bson_id).One(&game)
	})
	return game, err
}
