/**
 * Game controller to manage the game API
 */

 package controllers

 import (
	cb "licorne/app/controllers/base"
	"licorne/app/services/licorne_services"
	"licorne/utilities/tracelog"
	"licorne/app/models"
	"github.com/robfig/revel"
 )


type (
	// Game controller the buoy api
	GameController struct {
		cb.BaseController
	}
)

//** INIT FUNCTION

// init is called when the first request into the controller is made
func init() {
	revel.InterceptMethod((*GameController).Before, revel.BEFORE)
	revel.InterceptMethod((*GameController).After, revel.AFTER)
	revel.InterceptMethod((*GameController).Panic, revel.PANIC)
}

//** CONTROLLER FUNCTIONS

// Return list of all Games

// @Title getOrdersByCustomer
// @Description retrieves orders for given customer defined by customer ID
// @Accept  json
// @Param   customer_id     path    int     true        "Customer ID"
// @Param   order_id        query   int     false        "Retrieve order with given ID only"
// @Param   order_nr        query   string  false        "Retrieve order with given number only"
// @Param   created_from    query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created starting from created_from"
// @Param   created_to      query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created before created_to"
// @Success 200
// @Failure 400
// @Resource /games
// @Router /games [get]
func (this *GameController) Games() revel.Result {
	GamesList, err := licorne_services.GameAll(this.Services())
	if err != nil {
		return this.RenderText(err.Error())
	}

	return this.RenderJson(GamesList)
}

// Create a Game
func (this *GameController) CreateGame(name string) revel.Result {
	game := licorne_models.GameModel{name}
	tracelog.INFO(this.Services().UserId, "CreateGame", "game : %v", game)
	tracelog.INFO(this.Services().UserId, "CreateGame", "name : %v", name)
	tracelog.INFO(this.Services().UserId, "CreateGame", "params : %v", this.Params)
	err := licorne_services.CreateGame(this.Services(), game)
	if err != nil {
		return this.RenderText(err.Error())
	}

	return this.RenderJson(err)
}
