package tests

import(
	"github.com/robfig/revel"
	"encoding/json"
	"licorne/app/models"
	"licorne/utilities/helper"
)


type GamesTest struct {
	revel.TestSuite
}

func (t *GamesTest) Before() {
	println("Set up")
	revel.Config.SetSection("test")
	helper.MONGO_DATABASE = revel.Config.StringDefault("mgo.use_database", "")
}

func (t GamesTest) TestThatGameIndexWorks() {
	t.Get("/games")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	var a [0]licorne_models.GameModel
	body, _ := json.Marshal(a)
	println(revel.Config.StringDefault("mgo.use_database", ""))
	println(string(t.ResponseBody))
	t.Assert(string(t.ResponseBody) == string(body))
}

func (t *GamesTest) After() {
	println("Tear down")
}
