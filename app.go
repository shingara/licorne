package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"licorne/middleware"
	"licorne/controllers"
)

func main() {
	// mux := http.NewServeMux()


	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeHandler)

	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	db := GetDb(req)
	// 	site := getSite(db)
	// 	r.HTML(w, http.StatusOK, "home", site)
	// })

	n := negroni.Classic()
	n.Use(middleware.MongoMiddleware())
	n.UseHandler(router)
	n.Run(":9000")
}
