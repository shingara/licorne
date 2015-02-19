package controllers

import (
	"github.com/gin-gonic/gin"

	"licorne/services"
	"licorne/models"
)

func GamesIndexHandler(c *gin.Context) {
	games, _ := services.AllGame()
	json := services.MakeJsonGames(games)
	c.JSON(200, json)
}

func GamesCreateHandler(c *gin.Context) {
	var game models.GameForm
	c.Bind(&game)
	services.CreateGame(&game.Game)
	c.JSON(201, gin.H{"status": "success"})
}

