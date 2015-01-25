package middleware

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"gopkg.in/mgo.v2"
	"licorne/utilities"
	"os"
)

func MongoMiddleware() negroni.HandlerFunc {
	database := "licorne-dev"
	host := os.Getenv("DB_1_PORT_27017_TCP_ADDR")
	port := os.Getenv("DB_1_PORT_27017_TCP_PORT")

	session, err := mgo.Dial(host + ":" + port)

	if err != nil {
		panic(err)
	}

	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		reqSession := session.Clone()
		defer reqSession.Close()
		db := reqSession.DB(database)
		utilities.SetDb(r, db)
		next(rw, r)
	})
}

