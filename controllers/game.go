package controllers

import (
	"github.com/gin-gonic/gin"

	"licorne/services"
	"licorne/models"
)

// List all Games
func GamesIndexHandler(c *gin.Context) {
	games, _ := services.AllGame()
	json := models.MakeJsonGames(games)
	c.JSON(200, json)
}

// Create a Game
func GamesCreateHandler(c *gin.Context) {
	var game_form models.GameForm
	if c.Bind(&game_form) {
		game_model := models.ConvertToGame(game_form)
		services.CreateGame(&game_model)
		c.JSON(201, models.ConvertJsonGame(game_model))
	} else {
		c.JSON(400, c.Errors)
	}
}

// Get information of a Game
func GameShowHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	game, _ := services.GetGame(id)
	json := models.ConvertJsonGame(game)
	c.JSON(200, json)
}
