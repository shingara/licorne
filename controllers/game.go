package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"licorne/services"
	"licorne/models"
)

func GamesIndexHandler(c *gin.Context) {
	games, _ := services.AllGame()
	json := models.MakeJsonGames(games)
	c.JSON(200, json)
}

func GamesCreateHandler(c *gin.Context) {

	c.Request.ParseForm()
	fmt.Println(c.Request.Form)
	var game_form models.GameForm
	if c.Bind(&game_form) {
		game_model := models.ConvertToGame(game_form)
		services.CreateGame(&game_model)
		c.JSON(201, models.ConvertJsonGame(game_model))
	} else {
		c.JSON(400, c.Errors)
	}
}

