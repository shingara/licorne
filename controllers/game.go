package controllers

import (
	"net/http"
	"fmt"
	"licorne/services"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func GamesIndexHandler(rw http.ResponseWriter, r *http.Request) {
	games, _ := services.AllGame()
	json := services.MakeJsonGames(games)
	RenderJsonAPI(rw, 200, json)
}
