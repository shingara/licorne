package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"licorne/middleware"
	"licorne/controllers"
)


func BuildRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeHandler)
	router.HandleFunc("/games", controllers.GamesIndexHandler)
	return router
}

func main() {
	n := negroni.Classic()
	n.Use(middleware.MongoMiddleware())
	n.UseHandler(BuildRouter())
	n.Run(":9000")
}
