package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"licorne/controllers"
)

func BuildRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeHandler)
	router.HandleFunc("/v1/games", controllers.GamesIndexHandler)
	return router
}

func BuildMiddlewares() *negroni.Negroni {
	n := negroni.Classic()
	n.UseHandler(BuildRouter())
	return n
}

func main() {
	n := BuildMiddlewares()
	n.Run(":9000")
}
