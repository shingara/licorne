package utilities

import (
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"net/http"
)

func GetDb(r *http.Request) *mgo.Database {
	if rv := context.Get(r, "db"); rv != nil {
		return rv.(*mgo.Database)
	}
	return nil
}

func SetDb(r *http.Request, val *mgo.Database) {
	context.Set(r, "db", val)
}
