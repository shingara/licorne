package licorne_services

import (
	"licorne/app/models"
	"licorne/app/services"
	"licorne/utilities/tracelog"
	"licorne/utilities/helper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateGame(service *services.Service, game licorne_models.GameModel) (err error) {
	defer helper.CatchPanic(&err, service.UserId, "CreateGame")

	tracelog.STARTED(service.UserId, "CreateGame")
	tracelog.INFO(service.UserId, "CreateGame", "game : %v", game)

	err = service.DBAction("games",
		func(collection *mgo.Collection) error {
			return collection.Insert(game)
	})
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "CreateGame")
		return err
	}

	tracelog.COMPLETED(service.UserId, "CreateGame")
	return err
}
func GameAll(service *services.Service) (gameList []*licorne_models.GameModel, err error) {
	defer helper.CatchPanic(&err, service.UserId, "GameAll")

	tracelog.STARTED(service.UserId, "GameAll")

	//Find all games
	tracelog.TRACE(helper.MAIN_GO_ROUTINE, "GameAll", "Query : %s")

	// Capture the specified buoy
	gameList = []*licorne_models.GameModel{}
	err = service.DBAction("games",
		func(collection *mgo.Collection) error {
			return collection.Find(bson.M{}).All(&gameList)
		})

	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "GameAll")
		return gameList, err
	}

	tracelog.COMPLETED(service.UserId, "GameAll")
	return gameList, err
}
